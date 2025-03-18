package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS server
	nc, err := nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer nc.Close()

	// Create JetStream context
	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("Error creating JetStream context: %v", err)
	}

	streamName := "TASKS"
	consumerName := "worker"

	_, err = js.AddConsumer(streamName, &nats.ConsumerConfig{
		Durable:       consumerName,
		AckPolicy:     nats.AckExplicitPolicy,
		DeliverPolicy: nats.DeliverAllPolicy,
		MaxDeliver:    5,
		AckWait:       time.Second * 30,
		ReplayPolicy:  nats.ReplayInstantPolicy,
		FilterSubject: "tasks.*",
	})
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}

	sub, err := js.PullSubscribe("tasks.*", consumerName)
	if err != nil {
		log.Fatalf("Error subscribing to stream: %v", err)
	}

	// Wait for a termination signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	go func() {
		for {
			msgs, err := sub.Fetch(10, nats.MaxWait(time.Second*5))
			if err != nil {
				if err == nats.ErrTimeout {
					log.Printf("Warning: Timeout while fetching messages: %v", err)
					continue
				}
				log.Fatalf("Error fetching messages: %v", err)
			}
			for _, msg := range msgs {
				fmt.Printf("Received message: %s\n", string(msg.Data))
				msg.Ack()
			}
		}
	}()

	<-sigChan

	fmt.Println("Shutting down consumer service...")
	sub.Unsubscribe()
}
