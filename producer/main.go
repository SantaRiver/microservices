package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"sync"
	"time"
)

func main() {
	brokers := []string{"localhost:9092", "localhost:9093"}
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	syncProducer, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		log.Fatalf("sync kafka: %v", err)
	}
	asyncProducer, err := sarama.NewAsyncProducer(brokers, cfg)
	if err != nil {
		log.Fatalf("asyn kafka: %v", err)
	}

	go func() {
		for msg := range asyncProducer.Errors() {
			log.Printf("%v", msg)
		}
	}()
	wg := new(sync.WaitGroup)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(key string) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				asyncProducer.Input() <- &sarama.ProducerMessage{
					Topic: "user",
					Key:   sarama.StringEncoder(key),
					Value: sarama.ByteEncoder([]byte(fmt.Sprintf("%v -> %v", key, i))),
				}
			}
		}(fmt.Sprintf("%v", i))
	}
	par, off, err := syncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: "user",
		Key:   sarama.StringEncoder("sync"),
		Value: sarama.ByteEncoder([]byte(fmt.Sprintf("%v -> %v", "sync", time.Now()))),
	})
	log.Printf("%v -> %v; %v", par, off, err)
	wg.Wait()

}
