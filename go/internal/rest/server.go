package rest

import (
	"github.com/cleanarc/<% .ProjectName %>/internal/rest/handlers"
	"github.com/cleanarc/<% .ProjectName %>/internal/rest/routes"
	"github.com/cleanarc/<% .ProjectName %>/pkg/persistence"
	"github.com/cleanarc/go-core/rest"
	"log"
)

type Server struct {
	router     *rest.Router
	controller *handlers.Controller
}

func NewServer(db *persistence.DB) *Server {
	server := &Server{
		router:     rest.NewRouter(),
		controller: handlers.NewController(db),
	}

	server.setup()

	return server
}

func (s *Server) setup() {
	routes.Setup(s.router)
	// Docs & heath check endpoints
	routes.InitSwaggerRoutes(s.router)
	routes.InitStatusRoutes(s.router, s.controller)
	// TODO Add business domain endpoints here
}

func (s *Server) Run() {
	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	// router.Run(":3000") for a hard coded port
	err := s.router.Run()
	if err != nil {
		log.Panicf("failed to start %v", err)
	}
}
