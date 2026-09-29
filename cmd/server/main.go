package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"intensive.com/internal/account"
	"intensive.com/internal/auth"
	"intensive.com/internal/chat"
	"intensive.com/internal/database/generated"
)

func InitRouters(router *gin.Engine, q *generated.Queries) {

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

	accountHandler := account.New(q)
	accountHandler.RegisterRoutes(api)

	authHandler := auth.New(q)
	authHandler.RegisterRoutes(api)
}

func main() {
	router := gin.Default()

	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}
	queries := generated.New(db)

	InitRouters(router, queries)
	chat.ConfigureChatControllers(router)
	fmt.Println("Hello, world!")
	router.Run(os.Getenv("PORT"))
	fmt.Println("Hello, world!")
}
