package producer

import (
	"log"

	"github.com/IBM/sarama"
)

func NewProducer(topic string, data []byte) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	// Create a new Kafka producer
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Printf("Failed to create Kafka producer: %s\n", err)
		return err
	}
	defer producer.Close()

	// Create a Kafka message
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
		// Value: sarama.StringEncoder("Hello, Kafka!"),
	}

	// Send the message
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Println("Error Occurred:", err)
		return err
	}

	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	// log.Println("Message sent successfully!")
	return nil
}
