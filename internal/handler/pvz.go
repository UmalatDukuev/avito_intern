package handler

import (
	"avito_intern/internal/handler/entity"
	"avito_intern/internal/handler/request"
	"net/http"
	"strconv"
	"time"

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
	startDate := c.DefaultQuery("startDate", "")
	endDate := c.DefaultQuery("endDate", "")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	var start, end *time.Time
	if startDate != "" {
		parsedStart, err := time.Parse(time.RFC3339, startDate)
		if err == nil {
			start = &parsedStart
		}
	}
	if endDate != "" {
		parsedEnd, err := time.Parse(time.RFC3339, endDate)
		if err == nil {
			end = &parsedEnd
		}
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}

	pvzList, err := h.services.PVZ.GetPVZList(start, end, pageInt, limitInt)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pvz": pvzList,
	})
}
