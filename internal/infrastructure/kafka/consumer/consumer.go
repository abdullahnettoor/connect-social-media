package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)

func NewConsumer(topic, key string) {
	// Create a new Kafka consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %s", err)
	}
	defer consumer.Close()

	// Subscribe to the topic
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatalf("Failed to get partitions for topic '%s': %s", topic, err)
	}

	var wg sync.WaitGroup
	wg.Add(len(partitions))

	// Start a goroutine for each partition
	for _, partition := range partitions {
		go func(partition int32) {
			defer wg.Done()

			// Create a partition consumer
			partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
			if err != nil {
				log.Fatalf("Failed to create partition consumer: %s", err)
			}
			defer partitionConsumer.Close()

			// Process messages from the partition
			for message := range partitionConsumer.Messages() {
				log.Printf("Received message: %s", string(message.Value))
			}
		}(partition)
	}

	// Wait for a termination signal
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signals
		log.Println("Received interrupt signal. Shutting down...")
		os.Exit(0)
	}()

	// Wait for all partition consumers to finish
	wg.Wait()

}
