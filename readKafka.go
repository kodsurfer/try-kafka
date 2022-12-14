package try_kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Need a Kafka topic name")
		return
	}

	partition := 0
	topic := os.Args[1]
	fmt.Println("Kafka topic:", topic)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     topic,
		Partition: partition,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})
	r.SetOffset(0)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
		tmp := Record{}
		err = json.Unmarshal(m.Value, &tmp)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%T\n", tmp)
	}
	r.Close()
}
