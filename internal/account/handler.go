package account

import (
	"strconv"

	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type accountService interface {
	GetProfile(ctx context.Context, uid int32) (User, error)
}

type Handler struct {
	service accountService
}

func NewHandler(service accountService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetUserData(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	profile, err := h.service.GetProfile(c.Request.Context(), (int32)(uid))
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
