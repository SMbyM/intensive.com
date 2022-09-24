package controllers

import (
	"github.com/gin-gonic/gin"
	. "intensive.com/data"
	. "intensive.com/errors"
	. "intensive.com/models"
)

func Reg(c *gin.Context) {
	var user UserDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		go GlobalErrorsHandler.SendError(Error{
			Err:      err,
			Location: "package controllers >> authControllers.go >> line 12",
		})
		c.JSON(500, map[string]string{"error": err.Error()})
	}
	date := user.Day + "." + user.Month + "." + user.Year

	err := RegUser(user.Name, user.Lastname, user.Email, user.Password, user.Phone, user.Male, date)
	if err != nil {
		go GlobalErrorsHandler.SendError(Error{
			Err:      err,
			Location: "package controllers >> authControllers.go >> line 121",
		})
		c.HTML(500, "mistake.html", nil)
	}
	c.Redirect(301, "/account")
}

func LoginSite(c *gin.Context) {
	var user UserDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, map[string]string{"error": err.Error()})
	}

	err, _ := LoginUser(user.Name, user.Lastname, user.Email, user.Password)
	if err != nil {
		go GlobalErrorsHandler.SendError(Error{
			Err:      err,
			Location: "package controllers >> authControllers.go >> line 38",
		})
		c.HTML(500, "mistake.html", nil)
	}
	c.Redirect(301, "/account")
}

func GetUsersSearch(c *gin.Context) {
	users, err := GetUsers()

	if err != nil {
		go GlobalErrorsHandler.SendError(Error{
			Err:      err,
			Location: "package controllers >> authControllers.go >> line 50",
		})
		c.JSON(500, map[string]string{"error": err.Error()})
	}
	c.JSON(200, users)
}
