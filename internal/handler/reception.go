package handler

import (
	"avito_intern/internal/handler/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createReception(c *gin.Context) {
	var input request.CreateReception

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Reception.CreateReception(input.PvzID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) closeLastReception(c *gin.Context) {

}
