package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/petprojects9964409/todo_app/internal/models"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})

}

func (h *Handler) signIn(c *gin.Context) {

}
