package kafkahandler

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	kafkac "github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
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
	consumer, err := kafkac.NewConsumer(&kafkac.ConfigMap{
		"bootstrap.servers":    k.Host,
		"group.id":             "foo",
		"default.topic.config": kafkac.ConfigMap{"auto.offset.reset": "smallest"}})
	if err != nil {
		k.Logger.LogError("error initiate consumer : %s",err)
	}

	err = consumer.SubscribeTopics([]string{topic}, nil)
	run := true
	maxRetry := 0;
	for run == true {
		ev := consumer.Poll(500)
		switch e := ev.(type) {
		case *kafkac.Message:
			message = append(message, string(e.Value))
			fmt.Printf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))
		case kafkac.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
		case kafkac.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			maxRetry++
			fmt.Printf("Ignored %v\n", e)
		}

		if maxRetry == 10{
			run = false
		}
	}

	consumer.Close()

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

