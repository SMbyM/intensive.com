package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	. "intensive.com/controllers"
	. "intensive.com/data"
	. "intensive.com/errors"
)

func InitRouters(router *gin.Engine) {

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("intensive.comsession", store))

	router.LoadHTMLGlob("views/*.html")
	router.Static("/src", "./views/src")

	// router.POST("/login", LoginUserPost)

	router.POST("/intensive", nil)

	router.GET("/intensive", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "Intensive.html", nil)
	})

	router.GET("/account", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK,
			"account.html",
			gin.H{
				"name":     sessions.Default(ctx).Get("name"),
				"lastname": sessions.Default(ctx).Get("lastname"),
				"nickname": sessions.Default(ctx).Get("nickname"),
			})
	})

	router.GET("/users", GetUsersSearch)

	router.POST("/reg", Reg)

	router.POST("/login", LoginSite)

	router.POST("/set_friends", SetFriends)

	router.GET("/friends/:id", GetFriends)

	router.GET("/users_data/:id", GetUserData)

	router.GET("/get_search", GetUsersSearch)
}

func main() {
	router := gin.Default()

	db, err := sql.Open("sqlite3", "data/intensive.db")

	Db = db

	defer Db.Close()

	if err != nil {
		go GlobalErrorsHandler.SendError(Error{
			Err:      err,
			Location: "main.go >> line 62",
		})
		log.Fatal(err)
	}

	InitRouters(router)
	ConfigureChatControllers(router)

	router.Run(":8888")
}
