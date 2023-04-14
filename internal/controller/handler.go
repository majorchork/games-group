package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/majorchork/tech-crib-africa/config"
	"github.com/majorchork/tech-crib-africa/internal/models"
	"github.com/majorchork/tech-crib-africa/internal/port"
)

type Handler struct {
	DB     port.DB
	Config config.Config
}

func (h *Handler) Ping(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (h *Handler) GetUserFromContext(c *gin.Context) (*models.Admin, error) {
	authUser, exists := c.Get("user")
	if !exists {
		return nil, fmt.Errorf("error getting user from context")
	}
	user, ok := authUser.(*models.Admin)
	if !ok {
		return nil, fmt.Errorf("an error occurred")
	}
	return user, nil
}
