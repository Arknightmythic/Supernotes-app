package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateNotebookRequest struct {
	Name     string     `json:"name" validate:"required"`
	ParentId *uuid.UUID `json:"parent_id"`
}

type CreateNotebookResponse struct {
	Id uuid.UUID `json:"id"`
}

type ShowNotebookResponse struct {
	Id        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	ParentId  *uuid.UUID `json:"parent_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UpdateNotebookRequest struct {
	Id        uuid.UUID  `json:"id" validate:"required"`
	Name      string     `json:"name"`
}

type UpdateNotebookResponse struct {
	Id        uuid.UUID  `json:"id"`
}

type MoveNotebookRequest struct {
	Id        uuid.UUID  `json:"id"`
	ParentId  *uuid.UUID `json:"parent_id"`
}

type MoveNotebookResponse struct{
	Id uuid.UUID `json:"id"`
}

type GetAllNotebookResponses struct {
	Id        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	ParentId  *uuid.UUID `json:"parent_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}