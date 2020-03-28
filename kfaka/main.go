package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	//config
	config := sarama.NewConfig()
	//确认策略 leader follow 都需要确认
	config.Producer.RequiredAcks = sarama.WaitForAll
	//新选出一个partition
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true //成功交付的消息将在success channel 返回
	//构造一个消息
	message := &sarama.ProducerMessage{}
	message.Topic = "web_log"
	message.Value = sarama.StringEncoder("test log")
	//连接 kafak
	cli, err1 := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err1 != nil {
		fmt.Println("product close", err1)
		return
	}

	//defer cli.Close()
	//发送消息
	part, offset, err2 := cli.SendMessage(message)

	if err2 != nil {
		fmt.Println("send message failed", err2)
		return
	}
	fmt.Println(part, offset)
	err := cli.Close()
	if err != nil {
		fmt.Println("close failed", err)
	}
}
