package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Shopify/sarama"
	"log"
	client "ozon/internal/client/user"
	protos "ozon/internal/pkg/api"
	"ozon/saga/consts"
	"sync"
	"time"
)

type Map struct {
	data map[int]protos.User
	m    *sync.RWMutex
}

func (m *Map) Set(i int, u protos.User) {
	m.m.Lock()
	defer m.m.Unlock()
	m.data[i] = u
}

func (m *Map) Delete(i int) error {
	m.m.Lock()
	defer m.m.Unlock()
	if _, ok := m.data[i]; ok {
		delete(m.data, i)
		return nil
	}
	return errors.New("user not found")
}

func NewMap() *Map {
	return &Map{
		data: map[int]protos.User{},
		m:    new(sync.RWMutex),
	}
}

type User struct {
	data           *Map
	producer       sarama.SyncProducer
	incomeConsumer *IncomeHandler
}

type IncomeHandler struct {
	P    sarama.SyncProducer
	Data *Map
}

func (i IncomeHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (i IncomeHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (i IncomeHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Println("ConsumeClaim")
	for msg := range claim.Messages() {
		var user protos.User
		err := json.Unmarshal(msg.Value, &user)
		if err != nil {
			log.Printf("income data %v: %v", string(msg.Value), err)
			continue
		}

		i.Data.Set(int(user.ID), user)
		ctx := context.Background()

		_, err = client.CreateUser(ctx, &user)
		if err != nil {
			log.Printf("create user err: %v ", err)
			continue
		}
		log.Printf("create user")
	}
	return nil
}

func NewUser(ctx context.Context) (*User, error) {
	Data := NewMap()

	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(consts.Brokers, cfg)
	if err != nil {
		return nil, err
	}
	income, err := sarama.NewConsumerGroup(consts.Brokers, "user", cfg)
	if err != nil {
		return nil, err
	}
	iHandler := &IncomeHandler{
		P:    producer,
		Data: Data,
	}
	go func() {
		for {
			err := income.Consume(ctx, []string{"user"}, iHandler)
			if err != nil {
				log.Printf("income consumer error: %v", err)
			}
			time.Sleep(time.Second)
		}
	}()

	return &User{
		data:           NewMap(),
		producer:       producer,
		incomeConsumer: iHandler,
	}, nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	_, err := NewUser(ctx)
	if err != nil {
		log.Fatalf("NewUser: %v", err)
	}
	<-ctx.Done()
}
