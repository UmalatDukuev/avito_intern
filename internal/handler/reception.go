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
	pvz, err := h.services.PVZ.GetByID(input.PvzID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if pvz == nil {
		newErrorResponse(c, http.StatusNotFound, "there is no pvz with this id")
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
	pvzID := c.Param("pvzId")

	pvz, err := h.services.PVZ.GetByID(pvzID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if pvz == nil {
		newErrorResponse(c, http.StatusNotFound, "there is no pvz with this id")
		return
	}

	reception, err := h.services.Reception.CloseLastReception(pvzID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        reception.ID,
		"status":    reception.Status,
		"closed_at": reception.ClosedAt,
	})
}
