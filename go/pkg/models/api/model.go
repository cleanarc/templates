package api_models

import "github.com/google/uuid"

type Model struct {
	Id      uuid.UUID `json:"id"`
	Message string    `json:"message"`
}
