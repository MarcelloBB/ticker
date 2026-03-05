package controller

import (
	"errors"
	"net/http"

	"github.com/MarcelloBB/ticker/internal/dto"
	"github.com/MarcelloBB/ticker/internal/service"
	"github.com/gin-gonic/gin"
)

type UptimeController struct {
	service *service.UptimeService
}

func NewUptimeController(service *service.UptimeService) *UptimeController {
	return &UptimeController{service: service}
}

// CreateTarget godoc
// @Summary      Create uptime target
// @Description  Creates a new URL target for uptime monitoring
// @Tags         uptime
// @Accept       json
// @Produce      json
// @Param        payload  body      dto.CreateUptimeTargetRequest  true  "Target data"
// @Success      201      {object}  dto.UptimeTargetResponse
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /uptime/targets [post]
func (uc *UptimeController) CreateTarget(c *gin.Context) {
	var input dto.CreateUptimeTargetRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := uc.service.CreateTarget(c.Request.Context(), input)
	if err != nil {
		if errors.Is(err, service.ErrInvalidURL) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid URL, use full http:// or https:// URL"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

// ListTargets godoc
// @Summary      List uptime targets
// @Description  Lists all registered uptime targets
// @Tags         uptime
// @Produce      json
// @Success      200  {array}   dto.UptimeTargetResponse
// @Failure      500  {object}  map[string]string
// @Router       /uptime/targets [get]
func (uc *UptimeController) ListTargets(c *gin.Context) {
	result, err := uc.service.ListTargets(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
