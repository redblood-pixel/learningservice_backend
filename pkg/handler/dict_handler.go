package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redblood-pixel/learning-service-go/pkg/domain"
)

func (h *Handler) initDictRoutes(router *gin.RouterGroup) {
	dict := router.Group("/dict")
	{
		dict.GET("/", h.getAllWords)
		dict.GET("/:id", h.getWord)
		dict.POST("/", h.createWord)
		dict.PUT("/", h.updateWord)
		dict.DELETE("/", h.deleteWord)
	}
}

func (h *Handler) getAllWords(ctx *gin.Context) {

	res := h.service.Dictionary.GetAllWords()
	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) getWord(ctx *gin.Context) {

	wordIdString := ctx.Param("id")
	wordId, err := strconv.Atoi(wordIdString)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "word id must be possitive integer")
		return
	}
	fmt.Println(wordId)

	res, err := h.service.Dictionary.GetWord(wordId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusNotFound, "word not found")
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) createWord(ctx *gin.Context) {

	var (
		input domain.CreateWordRequest
		err   error
	)

	if err = ctx.ShouldBind(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "request is invalid")
		return
	}

	err = h.service.CreateWord(input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) updateWord(ctx *gin.Context) {
	var (
		input domain.Word
		err   error
	)

	if err = ctx.ShouldBind(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "request is invalid")
		return
	}

	err = h.service.UpdateWord(input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) deleteWord(ctx *gin.Context) {
	var (
		input domain.DeleteWordRequest
		err   error
	)

	if err = ctx.ShouldBind(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "request is invalid")
		return
	}

	err = h.service.DeleteWord(input.ID)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
}
