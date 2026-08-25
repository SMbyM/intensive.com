package account

import (
	"fmt"
	"strconv"

	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"intensive.com/internal/database"
)

type accountService interface {
	GetProfile(ctx context.Context, uid int) (UserProfile, error)
	SetFriendship(ctx context.Context, fst int, snd int) error
	GetFriendList(ctx context.Context, uid int) ([]UserProfile, error)
}

type Handler struct {
	service accountService
}

func NewHandler(service accountService) *Handler {
	return &Handler{service: service}
}

var Db = database.Db

func (h *Handler) GetFriendList(c *gin.Context) {

	uid, err := strconv.Atoi(c.Param("fst"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	friends, err := h.service.GetFriendList(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, friends)
}

func (h *Handler) SetFriendship(c *gin.Context) {
	var dto Friendship

	if err := c.ShouldBindJSON(&dto); err != nil {
		fmt.Println("accountControllers:18")
		c.HTML(505, "mistake.html", nil)
	}

	err := h.service.SetFriendship(c.Request.Context(), dto.Fst, dto.Snd)
	if err != nil {
		fmt.Println("accountControllers:24")
		c.HTML(505, "mistake.html", nil)
	}
}

func (h *Handler) GetUserData(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	profile, err := h.service.GetProfile(c.Request.Context(), uid)
	if err != nil {
		if errors.Is(err, ErrInvalidID) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user"})
		return
	}

	c.JSON(http.StatusOK, profile)
}
