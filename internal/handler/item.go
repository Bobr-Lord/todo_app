package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/petprojects9964409/todo_app/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(c *gin.Context) {
	userID, err := GetUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	var input models.TodoItem
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}

	id, err := h.service.TodoItem.Create(userID, listID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})

}

func (h *Handler) getAllItems(c *gin.Context) {
	userID, err := GetUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	items, err := h.service.TodoItem.GetAll(userID, listID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) getItemByID(c *gin.Context) {
	userID, err := GetUserID(c)
	if err != nil {
		return
	}

	itemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	item, err := h.service.TodoItem.GetByID(userID, itemID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)

}

func (h *Handler) updateItem(c *gin.Context) {
	userID, err := GetUserID(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	var input models.UpdateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.TodoItem.Update(userID, id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteItem(c *gin.Context) {
	userID, err := GetUserID(c)
	if err != nil {
		return
	}

	itemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.service.TodoItem.Delete(userID, itemID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
