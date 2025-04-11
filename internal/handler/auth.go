package handler

import (
	"avito_intern/internal/handler/request"
	"avito_intern/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) dummyLogin(c *gin.Context) {
	var input request.DummyLoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		println(err.Error())
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if input.UserType != "moderator" && input.UserType != "employee" {
		newErrorResponse(c, http.StatusInternalServerError, "user_type should be employee or moderator")
		return
	}
	token, err := h.services.Authorization.GenerateDummyToken(input.UserType)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to generate token")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) register(c *gin.Context) {
	var input request.RegisterInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if input.Role != "moderator" && input.Role != "employee" {
		newErrorResponse(c, http.StatusInternalServerError, "user_type should be employee or moderator")
		return
	}
	var registerInput service.RegisterInput
	registerInput.Email = input.Email
	registerInput.Password = input.Password
	registerInput.Role = input.Role

	id, err := h.services.Authorization.CreateUser(registerInput)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) login(c *gin.Context) {
	var input request.LoginInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var registerInput service.LoginInput
	registerInput.Email = input.Email
	registerInput.Password = input.Password

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
