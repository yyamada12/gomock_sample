package handler

import (
	"gin_sample/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewAdminHandler(service service.UserService) *AdminHandler {
	return &AdminHandler{
		service: service,
	}
}

type AdminHandler struct {
	service service.UserService
}

func (h AdminHandler) PostUser(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	// Parse JSON
	var json struct {
		Value string `json:"value" binding:"required"`
	}

	if c.Bind(&json) == nil {
		h.service.UpdateUser(user, json.Value)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
