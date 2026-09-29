package auth

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type authService interface {
	RegisterUser(ctx context.Context, user User) (User, error)
	LoginUser(ctx context.Context, user User) (User, error)
}

type Handler struct {
	service authService
}

func NewHandler(s authService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) RegisterUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("package controllers >> authControllers.go >> line 12")
		c.JSON(500, map[string]string{"error": err.Error()})
	}

	user, err := h.service.RegisterUser(c.Request.Context(), user)
	if err != nil {
		fmt.Println("package controllers >> authControllers.go >> line 121")
		c.HTML(500, "mistake.html", nil)
	}
	c.Redirect(301, "/account")
}

func (h *Handler) LoginSite(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, map[string]string{"error": err.Error()})
	}

	user, err := h.service.LoginUser(c.Request.Context(), user)
	if err != nil {
		fmt.Println("package controllers >> authControllers.go >> line 38")
		c.HTML(500, "mistake.html", nil)
	}
	c.Redirect(301, "/account")
}