package service

import (
	"avito_intern/internal/handler/entity"
	"avito_intern/internal/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "secretKey"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user RegisterInput) (string, error) {
	var us entity.User
	us.Email = user.Email
	us.Role = user.Role
	us.Password = user.Password
	return s.repo.CreateUser(us)
}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateDummyToken(userType string) (string, error) {
	claims := jwt.MapClaims{
		"userId":   "0",
		"userType": userType,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email, GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"userId": user.ID,
		"role":   user.Role,
		"exp":    time.Now().Add(tokenTTL).Unix(),
		"iat":    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (string, string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", errors.New("invalid token claims")
	}

	userIdStr, ok := claims["userId"].(string)
	if !ok {
		return "", "", errors.New("userId claim missing or invalid format")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "", "", errors.New("role claim missing")
	}

	return userIdStr, role, nil
}
