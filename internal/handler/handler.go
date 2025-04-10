package handler

import (
	"avito_intern/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.POST("/dummyLogin", h.dummyLogin)

	// auth := router.Group("/auth")
	// {
	// 	auth.POST("/register", h.register)
	// 	// auth.POST("/sign-in", h.signIn)

	// }
	return router
}
