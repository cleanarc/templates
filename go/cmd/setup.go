package main

import (
	"github.com/cleanarc/<% .ProjectName %>/docs"
	"github.com/cleanarc/<% .ProjectName %>/internal/rest"
	"github.com/cleanarc/<% .ProjectName %>/pkg/persistence"
)

func setup() *rest.Server {
	docs.Setup()

	db := persistence.Setup()
	server := rest.NewServer(db)

	return server
}
