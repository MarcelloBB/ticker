package controller

import (
	"net/http"

	"github.com/MarcelloBB/ticker/internal/dto"
	"github.com/gin-gonic/gin"
)

type HealthcheckController struct {}

func NewHealthcheckController() HealthcheckController {
	return HealthcheckController{}
}

// GetPing godoc
// @Summary      Ping healthcheck
// @Description  Returns pong
// @Produce      json
// @Success      200  {array}   dto.PingResponse
// @Failure      500  {object}  map[string]string
// @Router       /healthcheck [get]
func (uc *HealthcheckController) GetPing(c *gin.Context) {
	pong := dto.PingResponse{
		Message: "Pong",
	}
	c.JSON(http.StatusOK, pong)
}
