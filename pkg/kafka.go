package pkg

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	KafkaBatchMinBytes = 10 * 1024
	KafkaBatchMaxBytes = 1 * 1024 * 1024
	ReadTimeOut        = 10 * time.Second
	WriteTimeOut       = 10 * time.Second
)

func NewKafkaConnection(brokerAddr string, topic string, partition int) (*kafka.Conn, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", brokerAddr, topic, partition)
	if err != nil {
		return nil, err
	}
	// set write timeout
	conn.SetWriteDeadline(time.Now().Add(ReadTimeOut))

	// set read timeout
	conn.SetReadDeadline(time.Now().Add(WriteTimeOut))

	return conn, nil
}

func ProduceMessages(conn *kafka.Conn, values [][]byte) error {
	var messages []kafka.Message
	for _, val := range values {
		messages = append(messages, kafka.Message{Value: val})
	}

	_, err := conn.WriteMessages(messages...)
	if err != nil {
		return err
	}

	log.Printf("✅ Sent %d message(s)", len(messages))
	return nil
}

func ConsumeMessages(conn *kafka.Conn) ([][]byte, error) {
	batch := conn.ReadBatch(KafkaBatchMinBytes, KafkaBatchMaxBytes)
	defer batch.Close()

	var messages [][]byte
	b := make([]byte, 10*1024)

	for {
		n, err := batch.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("read error: %w", err)
		}
		msg := make([]byte, n)
		copy(msg, b[:n])
		messages = append(messages, msg)
	}

	return messages, nil
}
