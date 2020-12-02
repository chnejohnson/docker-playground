package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "Users", 0)
	if err != nil {
		log.Panic(err)
	}

	conn.SetDeadline(time.Now().Add(3 * time.Second))
	batch := conn.ReadBatch(10e3, 10e6) // 10KB, 10MB

	b := make([]byte, 10e6)

	for {
		n, err := batch.Read(b)
		if n == 0 {
			break
		}
		if err == io.EOF {
			fmt.Println("---end_of_Read---")
			break
		} else if err != nil {
			fmt.Println("Read batch error:", err)
			break
		}

		fmt.Println("Get Message: ", string(b))
	}

	// for {
	// 	msg, err := batch.ReadMessage()
	// 	if err != nil {
	// 		log.Println("Fail to read message:", err)
	// 		break
	// 	}
	// 	fmt.Println("GET Message:", string(msg.Value))
	// }

	if err := batch.Close(); err != nil {
		log.Panic("Fail to close batch: ", err)
	}

	if err := conn.Close(); err != nil {
		log.Panic("Fail to close kafka connection: ", err)
	}
}
