package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type DummyLoginInput struct {
	UserType string `json:"userType" binding:"required"`
}

var secretKey = []byte("secretKey")

func generateToken(userType string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userType"] = userType
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func (h *Handler) dummyLogin(c *gin.Context) {
	var input DummyLoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := generateToken(input.UserType)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to generate token")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) register(c *gin.Context) {
	// var input models.User

	// if err := c.BindJSON(&input); err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, "invalid input body")
	// 	return
	// }
	// id, err := h.services.Authorization.CreateUser(input)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, map[string]interface{}{
	// 	"id": id,
	// })
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": 2222,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	// var input signInInput

	// if err := c.BindJSON(&input); err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, map[string]interface{}{
	// 	"token": token,
	// })
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": 12312312312312312,
	})
}
