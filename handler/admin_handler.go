package handler

import (
	"gin_sample/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostUser(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	// Parse JSON
	var json struct {
		Value string `json:"value" binding:"required"`
	}

	if c.Bind(&json) == nil {
		service.UpdateUser(user, json.Value)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
