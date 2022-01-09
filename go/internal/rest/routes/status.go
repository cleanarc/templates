package routes

import (
	"github.com/cleanarc/<% .ProjectName %>/internal/rest/handlers"
	"github.com/cleanarc/go-core/rest"
)

func InitStatusRoutes(r *rest.Router, controller *handlers.Controller) {
	// Status endpoints
	r.Group(V1).GET("/status", controller.Status)
}
