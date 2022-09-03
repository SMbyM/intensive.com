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
)

func firstreg(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{
		"name":     ctx.PostForm("name"),
		"lastname": ctx.PostForm("lastname"),
		"nickname": ctx.PostForm("nickname"),
		"day":      ctx.PostForm("day"),
		"month":    ctx.PostForm("month"),
		"year":     ctx.PostForm("year"),
		"male":     ctx.PostForm("radio"),
	})
}

func InitRouters(router *gin.Engine) {

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("intensive.comsession", store))

	router.LoadHTMLGlob("views/*.html")
	router.Static("/src", "./views/src")

	// router.POST("/login", LoginUserPost)

	router.POST("/intensive", AddUsers)

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

	router.POST("/first_reg", firstreg)

	router.GET("/usersnames", GetUsersSearch)

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "search.prototype.html", nil)
	})

}

func main() {
	router := gin.Default()

	db, err := sql.Open("sqlite3", "data/intensive.db")

	Db = db

	defer Db.Close()

	if err != nil {
		log.Fatal(err)
	}

	InitRouters(router)
	ConfigureChatControllers(router)

	router.Run(":8888")
}
