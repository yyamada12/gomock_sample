package handler

import (
	"gin_sample/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

type UserHandler struct {
	service service.UserService
}

func (h *UserHandler) GetUser(c *gin.Context) {
	user := c.Params.ByName("name")
	value, ok := h.service.GetUser(user)
	if ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}
