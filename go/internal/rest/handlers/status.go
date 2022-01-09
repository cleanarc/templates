package handlers

import (
	api_models "github.com/cleanarc/<% .ProjectName %>/pkg/models/api"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// Status godoc
// @Summary Is the service alive?
// @Tags status
// @Accept  json
// @Produce  json
// @Success 200 {object} api_models.Model
// @Router /status [get]
func (c *Controller) Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, api_models.Model{
		Id:      uuid.New(),
		Message: "I'm Alive!",
	})
}
