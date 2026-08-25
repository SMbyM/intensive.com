package auth

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"intensive.com/internal/database"
)

type authService interface {
	RegUser(ctx context.Context, user UserProfile) error
	LoginUser(ctx context.Context, user UserProfile) error
}

type Handler struct {
	service authService
}

func NewHandler(s authService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Reg(c *gin.Context) {
	var user UserProfile
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("package controllers >> authControllers.go >> line 12")
		c.JSON(500, map[string]string{"error": err.Error()})
	}

	err := h.service.RegUser(c.Request.Context(), user)
	if err != nil {
		fmt.Println("package controllers >> authControllers.go >> line 121")
		c.HTML(500, "mistake.html", nil)
	}
	c.Redirect(301, "/account")
}

func (h *Handler) LoginSite(c *gin.Context) {
	var user UserProfile
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, map[string]string{"error": err.Error()})
	}

	err := h.service.LoginUser(c.Request.Context(), user)
	if err != nil {
		fmt.Println("package controllers >> authControllers.go >> line 38")
		c.HTML(500, "mistake.html", nil)
	}
	c.Redirect(301, "/account")
}

func GetUsersSearch(c *gin.Context) {
	users, err := database.GetUsers()

	if err != nil {
		fmt.Println("package controllers >> authControllers.go >> line 50")
		c.JSON(500, map[string]string{"error": err.Error()})
	}
	c.JSON(200, users)
}
