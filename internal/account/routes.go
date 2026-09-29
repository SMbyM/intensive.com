package account

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/users/:id", h.GetUserData)
}
