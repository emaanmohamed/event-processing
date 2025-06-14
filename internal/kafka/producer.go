package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

var producer *kafka.Producer

func InitProducer(broker string) {
	var err error
	_, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}
	log.Println("Kafka producer initialized")
}

func SendMessage(topic string, message []byte) error {
	deliveryChan := make(chan kafka.Event)

	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, deliveryChan)

	if err != nil {
		return err
	}

	event := <-deliveryChan
	msg := event.(*kafka.Message)

	if msg.TopicPartition.Error != nil {
		return msg.TopicPartition.Error
	}
	log.Printf("Message produced to topic %s [%d] at offset %d\n", *msg.TopicPartition.Topic, msg.TopicPartition.Partition, msg.TopicPartition.Offset)
	close(deliveryChan)
	return nil

}
