package kafkahandler

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

func(k *kafkahandler) Produce(topic string, partition int, message string)error{
	conn, err := kafka.DialLeader(context.Background(), k.Protocol, k.Host, topic, partition)
	if err != nil {
		k.Logger.LogError("failed to connect kafka %s because : %s", k.Host, err)
		return err
	}

	conn.SetWriteDeadline(time.Now().Add(10*time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(message)},
	)
	if err != nil {
		k.Logger.LogError("failed to write messages: %s", err)
		return err
	}

	if err = conn.Close(); err != nil {
		k.Logger.LogError("failed to close writer: %s", err)
		return err
	}
	return nil
}

func(k *kafkahandler) Consume(topic string, partition int)([]string, error){
	message := []string{}
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		k.Logger.LogError("failed to dial leader:", err)
		return message, err
	}

	conn.SetReadDeadline(time.Now().Add(10*time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		_, err = batch.Read(b)
		if err != nil {
			break
		}
		message = append(message, string(b))
		k.Logger.LogAccess("Consume kafka message %s:", string(b))
	}

	if err = batch.Close(); err != nil {
		k.Logger.LogError("failed to close batch:", err)
		return message, err
	}

	if err = conn.Close(); err != nil {
		k.Logger.LogError("failed to close connection:", err)
		return message, err
	}

	return message, nil
}


func(k *kafkahandler) Ping()error{
	conn, err := kafka.Dial(k.Protocol, k.Host)

	if err != nil {
		k.Logger.LogError("failed to connect kafka %s because : %s", k.Host, err)
		return err
	}

	if err = conn.Close(); err != nil {
		k.Logger.LogError("failed to close writer: %s", err)
		return err
	}
	return nil
}

