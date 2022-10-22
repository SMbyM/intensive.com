package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "intensive.com/data"
	. "intensive.com/models"
	"strconv"
)

func GetFriends(c *gin.Context) {
	friends := make(map[int]FriendsDTO)
	fst, err := strconv.Atoi(c.Param("fst"))
	if err != nil {
		fmt.Println("accountControllers:19")
		c.HTML(505, "mistake.html", nil)
	}
	r, err := Db.Query("select snd from friends where fst=$1", fst)
	if err != nil {
		fmt.Println("accountControllers:24")
		c.HTML(505, "mistake.html", nil)
	}

	for r.Next() {
		var snd int
		err = r.Scan(&snd)
		if err != nil {
			fmt.Println("accountControllers:31")
			c.HTML(505, "mistake.html", nil)
		}
		res, err := Db.Query("select name, lastname, nickname from users where id=$1", snd)
		if err != nil {
			fmt.Println("accountControllers:43")
			c.HTML(505, "mistake.html", nil)
		}
		for res.Next() {
			var friend FriendsDTO
			err := res.Scan(&friend.Name, &friend.Lastname, &friend.Nickname)
			if err != nil {
				fmt.Println("accountControllers:50")
				c.HTML(505, "mistake.html", nil)
			}

			friends[snd] = friend
		}
	}
	c.JSON(200, friends)
}

func SetFriends(c *gin.Context) {
	var dto FriendDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		fmt.Println("accountControllers:18")
		c.HTML(505, "mistake.html", nil)
	}

	_, err := Db.Exec("insert into friends values ($1, $2)", dto.Fst, dto.Snd)
	if err != nil {
		fmt.Println("accountControllers:24")
		c.HTML(505, "mistake.html", nil)
	}
}

func GetUserData(c *gin.Context) {
	var user UserDTO
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("accountControllers:79")
		c.HTML(505, "mistake.html", nil)
	}
	res, err := Db.Query("select * from users where id=$1", id)
	if err != nil {
		fmt.Println("accountControllers:84")
		c.HTML(505, "mistake.html", nil)
	}

	for res.Next() {
		err = res.Scan(&id, &user.Name, user.Lastname, &user.Nickname, &user.Email, &user.Phone, &user.Male, &user.Birthday)
		if err != nil {
			fmt.Println("accountControllers:91")
			c.HTML(505, "mistake.html", nil)
		}
	}
	c.JSON(200, user)
}
