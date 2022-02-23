package handlers

import (
	"github.com/KwokBy/easy-ops/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() UserHandler {
	return UserHandler{}
}

func (u *UserHandler) Login(c *gin.Context) {

	token, err := jwt.New(jwt.Data{UserID: 1})
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    gin.H{"token": token},
	})
}
