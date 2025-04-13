package handler

import (
	"avito_intern/internal/handler/entity"
	"avito_intern/internal/handler/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createPVZ(c *gin.Context) {
	var input request.CreatePVZInput

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	pvz := entity.PVZ{
		City: input.City,
	}

	id, err := h.services.PVZ.CreatePVZ(pvz)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) getPVZList(c *gin.Context) {

}
