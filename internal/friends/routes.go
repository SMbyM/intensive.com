package friends

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes(r *gin.RouterGroup) {
	r.POST("/create", h.CreateFriendship)
	r.POST("/delete", h.DeleteFriendship)
	r.GET("/:uid", h.GetFriendList)
	r.GET("/are_friends/:fst/:snd", h.AreFriends)
}