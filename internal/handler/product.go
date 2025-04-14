package handler

import (
	"avito_intern/internal/handler/request"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addProductToReception(c *gin.Context) {
	var input request.AddProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	validProductTypes := []string{"электроника", "одежда", "обувь"}
	if !contains(validProductTypes, input.Type) {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("invalid product type: %s, allowed values are: %v", input.Type, validProductTypes))
		return
	}

	id, err := h.services.Product.AddProductToReception(input.PvzID, input.Type)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func (h *Handler) deleteLastProduct(c *gin.Context) {
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

	productID, err := h.services.Product.DeleteLastProduct(pvzID)
	if err != nil {
		if err.Error() == "no active reception for this PVZ" {
			newErrorResponse(c, http.StatusBadRequest, "No active reception found for this PVZ")
			return
		}
		if err.Error() == "no products to delete" {
			newErrorResponse(c, http.StatusBadRequest, "No products to delete")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Product deleted successfully",
		"product_id": productID,
	})
}
