package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewErrorResponse(ctx *gin.Context, code int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(code, gin.H{
		"error": message,
	})
}
