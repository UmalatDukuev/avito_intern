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
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, role, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	
	c.Set(userCtx, userId)
	c.Set("role", role)
}

func (h *Handler) roleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, role, err := h.getUserRoleFromContext(c)
		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		if !isRoleAllowed(role, allowedRoles) {
			newErrorResponse(c, http.StatusForbidden, "access forbidden")
			return
		}

		c.Set(userCtx, userId)
		c.Set("role", role)
		c.Next()
	}
}

func isRoleAllowed(role string, allowedRoles []string) bool {
	for _, allowedRole := range allowedRoles {
		if role == allowedRole {
			return true
		}
	}
	return false
}

func (h *Handler) getUserRoleFromContext(c *gin.Context) (string, string, error) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return "", "", errors.New("user id not found")
	}

	userIdStr, ok := userId.(string)
	if !ok {
		return "", "", errors.New("user id is of invalid type")
	}

	role, ok := c.Get("role")
	if !ok {
		return "", "", errors.New("role not found")
	}

	roleStr, ok := role.(string)
	if !ok {
		return "", "", errors.New("role is of invalid type")
	}

	return userIdStr, roleStr, nil
}

func getUserId(c *gin.Context) (string, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return "0", errors.New("user id not found")
	}
	idStr, ok := id.(string)
	if !ok {
		return "0", errors.New("user id is of invalid type")
	}
	return idStr, nil
}
