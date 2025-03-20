package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.Request.Header.Get(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "No authorization header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header")
		return
	}
	userID, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header")
		return
	}
	c.Set(userCtx, userID)
}

func GetUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "No user id")
		return 0, errors.New("No user id")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "Invalid user id")
		return 0, errors.New("Invalid user id")
	}
	return idInt, nil
}
