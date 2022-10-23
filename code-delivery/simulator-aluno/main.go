package main

import (
	"fmt"

	kafka2 "github.com/fshreiner/Imersao-Full-Stack/tree/main/code-delivery/simulator-aluno/application/kafka"
	"github.com/fshreiner/Imersao-Full-Stack/tree/main/code-delivery/simulator-aluno/infra/kafka"

	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()
	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}
}
