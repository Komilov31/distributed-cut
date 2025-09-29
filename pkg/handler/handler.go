package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wb-go/wbf/zlog"
)

type Servcie interface {
	ProcessInput(input map[int]string) map[int]string
}

type Handler struct {
	service Servcie
}

func New(service Servcie) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) HandleInput(c *gin.Context) {
	var input map[int]string
	if err := c.BindJSON(&input); err != nil {
		zlog.Logger.Error().Msg("invalid payload: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	result := h.service.ProcessInput(input)
	c.JSON(http.StatusOK, result)
}
