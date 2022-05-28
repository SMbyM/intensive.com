package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"

	. "intensive.com/data"
)

func AddUsers(ctx *gin.Context) {
	err := RegUser(ctx.PostForm("name"), ctx.PostForm("lastname"), ctx.PostForm("email"), ctx.PostForm("password"))

	if err != nil {
		ctx.JSON(500, map[string]string{"error": err.Error()})
	}

	ctx.Redirect(301, "/regform")
}

func RegForm(ctx *gin.Context) {
	ctx.HTML(200, "regform.html", nil)
}

func GetUsersSearch(ctx *gin.Context) {
	users, err := GetUsers()

	if err != nil {
		ctx.JSON(500, map[string]string{"error": err.Error()})
	}
	ctx.JSON(200, users)
}

func DeleteCookies(ctx *gin.Context) {
	s := sessions.Default(ctx)

	s.Delete("name")
	s.Delete("email")
	s.Delete("password")

	s.Save()
}
