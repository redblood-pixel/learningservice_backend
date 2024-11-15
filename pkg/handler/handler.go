package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redblood-pixel/learning-service-go/internal/tokenutil"
	"github.com/redblood-pixel/learning-service-go/pkg/handler/middleware"
	"github.com/redblood-pixel/learning-service-go/pkg/service"
)

type Handler struct {
	service *service.Service
	tm      *tokenutil.TokenManager
}

func NewHandler(s *service.Service, tokenManager *tokenutil.TokenManager) *Handler {
	return &Handler{
		service: s,
		tm:      tokenManager,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()

	// Init users routes

	router.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "work",
		})
	})

	publicRouter := router.Group("")
	{
		h.initUsersRoutes(publicRouter)
	}

	protectedRouter := router.Group("/api", middleware.JwtAuthMiddleware(h.tm))
	{
		protectedRouter.GET("/ping", h.AuthPing)
		h.initDictRoutes(protectedRouter)
	}

	return router
}
