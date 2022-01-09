package handlers

import (
	"github.com/cleanarc/<% .ProjectName %>/pkg/persistence"
)

type Controller struct {
	db *persistence.DB
}

func NewController(db *persistence.DB) *Controller {
	return &Controller{
		db: db,
	}
}
