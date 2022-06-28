package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/ptypes/timestamp"
	"log"
	protos "ozon/internal/pkg/api"
	"ozon/pkg/randName"
	"ozon/saga/consts"
	"time"
)

func main() {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	syncProducer, err := sarama.NewSyncProducer(consts.Brokers, cfg)
	if err != nil {
		log.Fatalf("sync kafka: %v", err)
	}

	for {
		user := protos.User{
			ID:         int64(time.Now().UnixNano()),
			Name:       randName.New(),
			Surname:    randName.New(),
			Patronymic: randName.New(),
			Birthday:   &timestamp.Timestamp{Seconds: int64(time.Now().Unix())},
			Type:       1,
		}
		jUser, err := json.Marshal(&user)
		if err != nil {
			log.Printf("wtf? %v", err)
			continue
		}
		par, off, err := syncProducer.SendMessage(&sarama.ProducerMessage{
			Topic: "user",
			Key:   sarama.StringEncoder(fmt.Sprintf("%v", user.ID)),
			Value: sarama.ByteEncoder(jUser),
		})
		log.Printf("userID: %v \n %v -> %v; %v", user.ID, par, off, err)
		time.Sleep(time.Millisecond * 500)

		/*if rand.Intn(10) == 9 {
			par, off, err = syncProducer.SendMessage(&sarama.ProducerMessage{
				Topic: "reset_orders",
				Key:   sarama.StringEncoder(fmt.Sprintf("%v", d.Id)),
				Value: sarama.ByteEncoder(b),
			})
			log.Printf("reset %v -> %v; %v", par, off, err)
		}*/
	}
}
