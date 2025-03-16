package client

import (
	"fmt"
	"os"
)

func main() {
	natsUrl := os.Getenv("NATS_URL")
	fmt.Println(natsUrl)

}
