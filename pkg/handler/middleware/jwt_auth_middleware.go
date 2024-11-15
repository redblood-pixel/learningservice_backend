package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redblood-pixel/learning-service-go/internal/tokenutil"
	"github.com/sirupsen/logrus"
)

func JwtAuthMiddleware(tm *tokenutil.TokenManager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := ctx.Request.Header.Get("Authorization")
		ts := strings.Split(t, " ")
		if len(ts) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not valid token format",
			})
			return
		} else if ts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token must be Bearer",
			})
			return
		}
		id, err := tm.ParseAccessToken(ts[1])
		if err != nil {
			logrus.Error(err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.Set("user_id", id)
		ctx.Next()
	}
}
