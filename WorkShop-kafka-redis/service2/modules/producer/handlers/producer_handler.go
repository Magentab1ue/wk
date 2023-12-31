package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2/log"

	"service2/modules/entities/events"
	"service2/modules/entities/models"
)

type eventProducer struct {
	producer sarama.SyncProducer
}

func NewEventHandlerProducer(producer sarama.SyncProducer) models.EventHandlerProduce {
	return &eventProducer{producer}
}

func (obj *eventProducer) Produce(event events.Event) error {
	topic := event.String()
	value, err := json.Marshal(event)
	if err != nil {
		log.Error(err)
		return err
	}

	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}

	p, o, err := obj.producer.SendMessage(&msg)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info(fmt.Sprintf("sent to topic: %v, partition: %v, offset %v", topic, p, o))
	return nil
}
