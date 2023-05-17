```go
package main

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"math/rand"
	"time"
)

func main() {
	channel := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)

	// Subscribe 1
	messages, err := channel.Subscribe(context.TODO(), "example.topic")
	if err != nil {
		fmt.Println(err)
	}
	go process1(messages)

	// Subscribe 2
	messages2, err := channel.Subscribe(context.TODO(), "example.topic2")
	if err != nil {
		fmt.Println(err)
	}
	go process2(messages2)

	// publish
	publishMessages(channel)
}

func publishMessages(channel message.Publisher) {
	time.Sleep(4 * time.Second)
	for {
		n := rand.Int()
		msg := message.NewMessage(watermill.NewUUID(), []byte(fmt.Sprintf("Hello, world! %d", n)))

		topic := "example.topic"
		if n%2 == 0 {
			topic = "example.topic2"
		}

		if err := channel.Publish(topic, msg); err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Second)
	}
}

func process1(messages <-chan *message.Message) {
	fmt.Println("process1 is ready")
	for msg := range messages {
		fmt.Printf("process1 received message: %s, payload: %s\n", msg.UUID, string(msg.Payload))

		msg.Ack()
	}
}

func process2(messages <-chan *message.Message) {
	fmt.Println("process2 is ready")
	for msg := range messages {
		fmt.Printf("process2 received message: %s, payload: %s\n", msg.UUID, string(msg.Payload))

		msg.Ack()
	}
}
```
