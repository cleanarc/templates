package db_models

import "github.com/google/uuid"

type Model struct {
	Id uuid.UUID `gorm:"unique;type:uuid"`
}
