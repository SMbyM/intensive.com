package friends

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"net/http"
)


type friendsService interface {
	CreateFriendship(ctx context.Context, fst int32, snd int32) error
	DeleteFriendship(ctx context.Context, fst int32, snd int32) error
	GetFriendList(ctx context.Context, uid int32) ([]int32, error)
	AreFriends(ctx context.Context, fst int32, snd int32) (bool, error)
}

type Handler struct {
	service friendsService
}

func NewHandler(service friendsService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateFriendship(c *gin.Context) {
	var dto FriendDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.CreateFriendship(c.Request.Context(), (int32)(dto.Fst), (int32)(dto.Snd))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "friendship created"})
}

func (h *Handler) DeleteFriendship(c *gin.Context) {
	var dto FriendDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.DeleteFriendship(c.Request.Context(), (int32)(dto.Fst), (int32)(dto.Snd))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "friendship deleted"})
}

func (h *Handler) GetFriendList(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	friends, err := h.service.GetFriendList(c.Request.Context(), (int32)(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, friends)
}

func (h *Handler) AreFriends(c *gin.Context) {
	fst, err := strconv.Atoi(c.Param("fst"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid first id"})
		return
	}

	snd, err := strconv.Atoi(c.Param("snd"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid second id"})
		return
	}

	areFriends, err := h.service.AreFriends(c.Request.Context(), (int32)(fst), (int32)(snd))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"are_friends": areFriends})
}