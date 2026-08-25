package account

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/users/:id", h.GetUserData)
	r.POST("/set_friends", h.SetFriendship)
	r.GET("/friends/:id", h.GetFriendList)
}
