package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/Shopify/sarama"
)

// exampleConsumerGroupHandler 实现了 sarama.ConsumerGroupHandler 接口
type exampleConsumerGroupHandler struct{}

// Setup 在消费者组启动之前调用，用于初始化一些资源
func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

// Cleanup 在消费者组停止之后调用，用于清理一些资源
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim 在每个分区上执行消费逻辑
func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// 从 claim.Messages() 通道中循环读取消息
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		// 标记消息已处理，sarama会自动提交
		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	// 创建一个新的消费者组
	group, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "test-group", nil)
	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()

	// 设置一个信号处理器，用于优雅地退出
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := group.Consume(ctx, []string{"test"}, exampleConsumerGroupHandler{}); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err = group.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}
