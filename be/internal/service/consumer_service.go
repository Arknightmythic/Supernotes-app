package service

import (
	"ai-notetaking-be/internal/dto"
	"ai-notetaking-be/internal/entity"
	"ai-notetaking-be/internal/repository"
	"ai-notetaking-be/pkg/embedding"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

type IConsumerService interface {
	Consume(ctx context.Context) error
}

type consumerService struct {
	notebookRepository      repository.INotebookRepository
	noteRepository          repository.INoteRepository
	noteEmbeddingRepository repository.INoteEmbeddingRepository
	pubSub                  *gochannel.GoChannel
	topicName               string
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

	note, err := cs.noteRepository.GetById(ctx, payload.NoteId)
	if err != nil {
		panic(err)
	}

	
	notebook, err := cs.notebookRepository.GetById(ctx, note.NotebookId )
	if err != nil {
		panic(err)
	}

	noteUpdateAt := "-"
	if note.UpdatedAt != nil {
		noteUpdateAt = note.UpdatedAt.Format(time.RFC3339)
	}

	content := fmt.Sprintf(`
	Note Title: %s
	Notebook Title: %s
	%s

	Created At: %s
	Updated At: %s
	`,
		note.Title, notebook.Name, note.Content, note.CreatedAt.Format(time.RFC3339), noteUpdateAt,
	)

	res, err := embedding.GetGeminiEmbedding(os.Getenv("GOOGLE_GEMINI_API_KEY"), note.Content)
	if err != nil {
		panic(err)
	}


	
	
	noteEmbedding := entity.NoteEmbedding{
		Id:             uuid.New(),
		Document:       content,
		EmbeddingValue: res.Embbeding.Values,
		NoteId:         note.Id,
		CreatedAt:      time.Now(),
	}

	err = cs.noteEmbeddingRepository.Create(ctx, &noteEmbedding)

	if err != nil {
		panic(err)
	}

	msg.Ack()
}

func NewConsumerService(pubSub *gochannel.GoChannel, topicName string, noteRepository repository.INoteRepository, noteEmbeddingRepository repository.INoteEmbeddingRepository, notebookRepository repository.INotebookRepository) IConsumerService {
	return &consumerService{
		pubSub:                  pubSub,
		topicName:               topicName,
		noteRepository:          noteRepository,
		noteEmbeddingRepository: noteEmbeddingRepository,
		notebookRepository:      notebookRepository,
	}
}
