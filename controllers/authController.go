package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	. "intensive.com/data"
)

type UserDTO struct {
}

func AddUsers(c *gin.Context) {
	var user UserDTO
	err := c.ShouldBind(&user)

	if err != nil {
		c.JSON(500, map[string]string{"error": err.Error()})
	}

	s := sessions.Default(c)

	s.Set("name", c.PostForm("name"))
	s.Set("lastname", c.PostForm("lastname"))
	s.Set("email", c.PostForm("email"))

	c.Redirect(301, "/account")
}

func GetUsersSearch(c *gin.Context) {
	users, err := GetUsers()

	if err != nil {
		c.JSON(500, map[string]string{"error": err.Error()})
	}
	c.JSON(200, users)
}
