package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redblood-pixel/learning-service-go/pkg/domain"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {

	users := api.Group("/users")
	{
		users.POST("/sign-up", h.userSignup)
		users.POST("/sign-in", h.userSignin)
		users.POST("/refresh", h.userRefreshToken)
		// TODO logout handler
	}
}

// TODO add error status code identifier

func (h *Handler) userSignup(ctx *gin.Context) {

	var input domain.SignupInput

	if err := ctx.ShouldBind(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.service.Users.SignUp(input)

	if err != nil {
		//
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, tokens)
}

func (h *Handler) userSignin(ctx *gin.Context) {
	var input domain.SigninInput

	if err := ctx.ShouldBind(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.service.Users.SignIn(input)

	if err != nil {
		//
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, tokens)
}

func (h *Handler) userRefreshToken(ctx *gin.Context) {

	var input domain.RefreshInput

	if err := ctx.ShouldBind(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.service.Users.Refresh(input)

	if err != nil {
		//
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, tokens)
}

func (h *Handler) AuthPing(ctx *gin.Context) {
	id, ok := ctx.Get("user_id")
	if !ok {
		NewErrorResponse(ctx, http.StatusUnauthorized, "No userid in context.")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "work",
		"id":      id,
	})
}
