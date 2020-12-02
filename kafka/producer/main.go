package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

const msg = "Grace"

func main() {

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "Users", 0)
	if err != nil {
		log.Panic("Fail to Dial Leader: ", err)
	}

	_, err = conn.WriteMessages(kafka.Message{Value: []byte(msg)})
	if err != nil {
		log.Panic("Fail to write message: ", err)
	}

	if err := conn.Close(); err != nil {
		log.Panic("Fail to close connection: ", err)
	}

	fmt.Printf("Succeed to produce message %s\n", msg)
}
