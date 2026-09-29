package auth

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/reg", h.RegisterUser)
	r.POST("/login", h.LoginSite)
}
