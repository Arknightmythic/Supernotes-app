package service

import (
	"ai-notetaking-be/internal/dto"
	"ai-notetaking-be/internal/repository"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/gofiber/fiber/v2/log"
)



type IConsumerService interface {
	Consume(ctx context.Context) error
}

type consumerService struct {
	noteRepository repository.INoteRepository
	pubSub         *gochannel.GoChannel
	topicName      string
}


func (cs *consumerService) Consume(ctx context.Context) error {
	messages, err := cs.pubSub.Subscribe(ctx, cs.topicName)
	if err != nil {
		return err
	}

	go func() {
		for msg := range messages {
			cs.ProcessMessage(ctx, msg)
		}
	}()

	return nil
}

func (cs *consumerService) ProcessMessage(ctx context.Context, msg *message.Message) {

	defer msg.Ack()
	defer func() {
		if e := recover(); e != nil {
			log.Error(e)
		}
	}()

	var payload dto.PublishEmbedNoteMessage
	err := json.Unmarshal(msg.Payload, &payload)
	if err != nil {
		panic(err)
	}

	Note, err := cs.noteRepository.GetById(ctx, payload.NoteId)
	if err != nil {
		panic(err)
	}

	

	fmt.Println(resEmbedding.Embbeding.Values)
	msg.Ack()
}

func NewConsumerService(pubSub *gochannel.GoChannel, topicName string, noteRepository repository.INoteRepository) IConsumerService {
	return &consumerService{
		pubSub:         pubSub,
		topicName:      topicName,
		noteRepository: noteRepository,
	}
}
