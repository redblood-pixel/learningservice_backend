package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redblood-pixel/learning-service-go/pkg/domain"
)

func (h *Handler) initGroupRoutes(router *gin.RouterGroup) {
	group := router.Group("/groups")
	{
		// Common
		group.POST("/", h.createGroup)
		group.GET("/all", h.getAllGroups)
		group.GET("/:id", h.getGroup)
		group.PUT("/", h.updateGroup)
		group.DELETE("/", h.deleteGroup)

		// Important routes
		group.GET("/words/:id", h.getWordsInGroup)
		group.GET("/my", h.getGroupsOfUser)
	}
}

func (h *Handler) createGroup(c *gin.Context) {
	var input domain.CreateGroupRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	id, err := h.service.Group.CreateGroup(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) getAllGroups(c *gin.Context) {

	res := h.service.Group.GetAllGroups()

	c.JSON(http.StatusOK, res)
}

func (h *Handler) getGroup(c *gin.Context) {
	groupIDString := c.Param("id")
	groupID, err := strconv.Atoi(groupIDString)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "word id must be possitive integer")
		return
	}

	res, err := h.service.Group.GetGroup(groupID)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) deleteGroup(c *gin.Context) {
	var input domain.DeleteGroupRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	err := h.service.Group.DeleteGroup(input.ID)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, struct{}{})
}

func (h *Handler) updateGroup(c *gin.Context) {
	var input domain.Group

	if err := c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	err := h.service.Group.UpdateGroup(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, struct{}{})
}

func (h *Handler) getWordsInGroup(c *gin.Context) {
	groupIDString := c.Param("id")
	groupID, err := strconv.Atoi(groupIDString)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "word id must be possitive integer")
		return
	}

	res, err := h.service.Group.GetWordsInGroup(groupID)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) getGroupsOfUser(c *gin.Context) {
	userID := c.GetInt("user_id")

	res, err := h.service.Group.GetGroupsOfUser(userID)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, res)
}
