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
	router.POST("/register", h.register)
	router.POST("/login", h.login)
	pvz := router.Group("/pvz")
	{
		pvz.POST("/", h.userIdentity, h.roleMiddleware("moderator"), h.createPVZ)
		pvz.GET("/", h.getPVZList)
		pvzID := router.Group("/:pvzId")
		{
			pvzID.POST("/close_last_reception", h.closeLastReception)
			pvzID.POST("/delete_last_product", h.deleteLastProduct)
		}
	}
	router.POST("/receptions", h.userIdentity, h.roleMiddleware("employee"), h.createReception)
	router.POST("/products", h.addProductToReception)
	return router
}
