package service

import (
	"ai-notetaking-be/internal/dto"
	"context"
	"encoding/json"
	"fmt"

	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

type IConsumerService interface {
	Consume(ctx context.Context) error
}

type consumerService struct {
	pubSub    *gochannel.GoChannel
	topicName string
}

func (cs *consumerService) Consume(ctx context.Context) error {
	messages, err := cs.pubSub.Subscribe(ctx, cs.topicName)
	if err != nil {
		return err
	}

	go func() {
		for msg := range messages {
			var payload dto.PublishEmbedNoteMessage
			err := json.Unmarshal(msg.Payload, &payload)

			if err != nil {
				panic(err)
			}

			fmt.Println(payload.NoteId)
			msg.Ack()
		}
	}()

	return nil
}

func NewConsumerService(pubSub *gochannel.GoChannel, topicName string) IConsumerService {
	return &consumerService{
		pubSub:    pubSub,
		topicName: topicName,
	}
}
