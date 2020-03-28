package kfaka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var cli sarama.SyncProducer

func Init(addr []string) (err error) {
	//配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	cli, err = sarama.NewSyncProducer(addr, config)
	return
}

func SendToKfaka(topic, msg string) {
	msgObj := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}
	_, _, err := cli.SendMessage(msgObj)
	if err != nil {
		fmt.Println("send message to kafka failed", err)
	}
	//fmt.Println(part,offset)
}
