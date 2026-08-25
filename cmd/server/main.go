package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"intensive.com/internal/account"
	"intensive.com/internal/auth"
	"intensive.com/internal/chat"
	"intensive.com/internal/database"
)

func InitRouters(router *gin.Engine) {

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("intensive.comsession", store))

	router.LoadHTMLGlob("web/views/*.html")
	router.Static("/src", "web/views/src")

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
	api := router.Group("/api")

	accountHandler := account.New(database.Db)
	accountHandler.RegisterRoutes(api)

	authHandler := auth.New(database.Db)
	authHandler.RegisterRoutes(api)

	router.GET("/users", auth.GetUsersSearch)

	router.GET("/get_search", auth.GetUsersSearch)
}

func main() {
	router := gin.Default()

	db, err := database.Connect(os.Getenv("DATABASE_URL"))
	database.Db = db
	if err != nil {
		panic(err.Error())
	}
	defer database.Db.Close()

	InitRouters(router)
	chat.ConfigureChatControllers(router)
	fmt.Println("Hello, world!")
	router.Run(":8080")
	fmt.Println("Hello, world!")
}
