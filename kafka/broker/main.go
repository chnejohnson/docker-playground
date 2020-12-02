package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.DialContext(context.Background(), "tcp", "localhost:9092")
	if err != nil {
		log.Panic("Fail to connect to kafka: ", err)
	}
	defer conn.Close()
	fmt.Println("kafka connected!")

	topicConfigs := kafka.TopicConfig{
		Topic:             "Users",
		NumPartitions:     1,
		ReplicationFactor: 1,
	}

	if err := conn.CreateTopics(topicConfigs); err != nil {
		log.Panic("Fail to create topic: ", err)
	}

	fmt.Printf("Success to create topic of %+v\n", topicConfigs)

}
