package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS server
	fmt.Printf("NATS_URL: %s\n", os.Getenv("NATS_URL"))
	nc, err := nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		fmt.Println("Error connecting to NATS server")
		log.Fatal(err)
	}
	defer nc.Close()

	// Create JetStream context
	js, err := nc.JetStream()
	if err != nil {
		fmt.Println("Error creating JetStream context")
		log.Fatal(err)
	}

	//Create a stream with workqueue policie and 1GB storage limit
	streamName := "TASKS"
	_, err = js.AddStream(&nats.StreamConfig{
		Name:      streamName,
		Subjects:  []string{"tasks.>"},
		Storage:   nats.FileStorage,
		Retention: nats.WorkQueuePolicy,
		MaxAge:    24 * time.Hour,
	})
	if err != nil {
		fmt.Println("Error creating stream")
		log.Fatal(err)
	}

	// Publish messages to the subject
	for i := range 10000 {
		lowRate, mediumRate := 0.6, 0.2
		rate := rand.Float64()
		var priority string
		if rate < lowRate {
			priority = "low"
		} else if rate < lowRate+mediumRate {
			priority = "medium"
		} else {
			priority = "high"
		}
		msg := &nats.Msg{
			Subject: "tasks." + priority,
			Data:    []byte(fmt.Sprintf("[%v] Message: %d", priority, i)),
		}
		_, err = js.PublishMsg(msg)
		if err != nil {
			fmt.Println("Error publishing message")
			log.Fatal(err)
		}
	}

}
