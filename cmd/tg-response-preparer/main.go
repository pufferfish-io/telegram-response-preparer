package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	cfg "tg-response-preparer/internal/config"
	cfgModel "tg-response-preparer/internal/config/model"
	"tg-response-preparer/internal/messaging"
	"tg-response-preparer/internal/processor"
	"tg-response-preparer/internal/service"
)

func main() {
	log.Println("🚀 Starting tg-response-preparer...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	kafkaConf, err := cfg.LoadSection[cfgModel.KafkaConfig]()
	if err != nil {
		log.Fatalf("❌ Failed to load Kafka config: %v", err)
	}

	if err != nil {
		log.Fatalf("❌ Failed to create S3 uploader: %v", err)
	}

	convertor := service.NewMessageConvertor()
	tgMessagePreparer := processor.NewTgMessagePreparer(kafkaConf.TelegramMessageTopicName, convertor)
	handler := func(msg []byte) error {
		return tgMessagePreparer.Handle(ctx, msg)
	}

	messaging.Init(kafkaConf.BootstrapServersValue)

	consumer, err := messaging.NewConsumer(kafkaConf.BootstrapServersValue, kafkaConf.ResponseMessageGroupId, kafkaConf.ResponseMessageTopicName, handler)
	if err != nil {
		log.Fatalf("❌ Failed to create consumer: %v", err)
	}

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		cancel()
	}()

	if err := consumer.Start(ctx); err != nil {
		log.Fatalf("❌ Consumer error: %v", err)
	}
}
