package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"gopkg.in/olahol/melody.v1"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Msg struct {
	User int64  `json:"user"`
	Data string `json:"text"`
	Chat int64  `json:"chat"`
}
type MsgDTO struct {
	User string `json:"user"`
	Data string `json:"text"`
	Chat string `json:"chat"`
}

type DTO struct {
	name string `json:"name" binding:"require"`
}

func ConfigureChatControllers(r *gin.Engine) {
	m := melody.New()

	r.Static("/static", "./static")

	r.LoadHTMLGlob("*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/channel/:name", func(c *gin.Context) {
		var (
			chatList []string
		)
		DB, err := sql.Open("sqlite3", "data/intensive.db")
		if err != nil {
			c.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
			return
		}
		r, err := DB.Query("select name from chat")
		if err != nil {
			c.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
			return
		}
		for r.Next() {
			var name string
			_ = r.Scan(&name)
			chatList = append(chatList, name)
		}
		name := c.Param("name")
		allowedInDb := slices.IndexFunc(chatList, func(e string) bool {
			return name == e
		})
		if allowedInDb != -1 {
			c.HTML(http.StatusOK, "chan.html", nil)
			return
		} else {
			c.JSON(http.StatusNotFound, map[string]string{
				"error": "Chat name not allowed in Db",
			})
			return
		}
	})

	r.GET("/id/:mod/:name", func(c *gin.Context) {
		if c.Param("mod") == "user" {
			var dto DTO
			var uid int
			dto.name = c.Param("name")
			db, err := sql.Open("sqlite3", "data/intensive.db")
			if err != nil {
				fmt.Print("Ошибка на 92 строке:")
				fmt.Println(err.Error())
			}
			r, err := db.Query("select id from users where name = $1", dto.name)
			if err != nil {
				fmt.Print("Ошибка на 97 строке:")
				fmt.Println(err.Error())
			}
			for r.Next() {
				err = r.Scan(&uid)
				if err != nil {
					fmt.Print("Ошибка на 103 строке:")
					fmt.Println(err.Error())
				}
			}
			err = db.Close()
			if err != nil {
				fmt.Print("Ошибка на 106 строке:")
				fmt.Println(err.Error())
			}
			c.JSON(200, map[string]any{
				"id": uid,
			})
		} else if c.Param("mod") == "chat" {
			var dto DTO
			var cid int
			dto.name = c.Param("name")
			db, err := sql.Open("sqlite3", "data/intensive.db")
			if err != nil {
				fmt.Print("Ошибка на 121 строке:")
				fmt.Println(err.Error())
			}
			r, err := db.Query("select id from chat where name = $1", dto.name)
			if err != nil {
				fmt.Print("Ошибка на 126 строке:")
				fmt.Println(err.Error())
			}
			for r.Next() {
				err = r.Scan(&cid)
				if err != nil {
					fmt.Print("Ошибка на 132 строке:")
					fmt.Println(err.Error())
				}
				fmt.Println(cid)
			}
			err = db.Close()
			if err != nil {
				fmt.Print("Ошибка на 139 строке:")
				fmt.Println(err.Error())
			}
			c.JSON(200, map[string]any{
				"id": cid,
			})
		} else {

		}
	})

	r.GET("/channel/:name/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	r.GET("/channel/massages/:chatId/:period", func(c *gin.Context) {
		db, err := sql.Open("sqlite3", "data/intensive.db")
		if err != nil {
			fmt.Print("Ошибка на 157 строке:")
			fmt.Println(err.Error())
		}
		response := make(map[string]MsgDTO)
		r, err := db.Query("select from_, message from messages where chat_id = $1", c.Param("chatId"))
		if err != nil {
			fmt.Print("Ошибка на 163 строке:")
			fmt.Println(err.Error())
		}
		i := 1
		for r.Next() {
			var (
				msg    MsgDTO
				userId int
			)
			err = r.Scan(&userId, &msg.Data)
			if err != nil {
				fmt.Print("Ошибка на 174 строке:")
				fmt.Println(err.Error())
			}
			re, err := db.Query("select name from users where id = $1", userId)
			if err != nil {
				fmt.Print("Ошибка на 179 строке:")
				fmt.Println(err.Error())
			}
			for re.Next() {
				_ = re.Scan(&msg.User)
			}
			response[string(i)] = msg
			fmt.Println(response[string(i)])
			i = i + 1
		}
		err = db.Close()
		if err != nil {
			fmt.Print("Ошибка на 191 строке:")
			fmt.Println(err.Error())
		}
		c.JSON(200, response)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		var dto MsgDTO
		var ms Msg
		_ = json.Unmarshal(msg, &dto)
		ms.Chat, _ = strconv.ParseInt(dto.Chat, 10, 0)
		ms.User, _ = strconv.ParseInt(dto.User, 10, 0)
		ms.Data = dto.Data
		fmt.Println(ms)

		d, err := sql.Open("sqlite3", "data/intensive.db")
		if err != nil {
			fmt.Print("Ошибка на 208 строке:")
			fmt.Println(err.Error())
			return
		}
		r, err := d.Query("select name from chat where id = $1", ms.Chat)
		if err != nil {
			fmt.Print("Ошибка на 214 строке:")
			fmt.Println(err.Error())
			return
		}
		re, err := d.Query("select name from users where id = $1", ms.User)
		if err != nil {
			fmt.Print("Ошибка на 220 строке:")
			fmt.Println(err.Error())
			return
		}
		for r.Next() && re.Next() {
			err = r.Scan(&dto.Chat)
			if err != nil {
				fmt.Print("Ошибка на 227 строке:")
				fmt.Println(err.Error())
			}
			err = re.Scan(&dto.User)
			if err != nil {
				fmt.Print("Ошибка на 232 строке:")
				fmt.Println(err.Error())
			}
		}
		r.Close()
		re.Close()
		msg, err = json.Marshal(dto)
		if err != nil {
			fmt.Print("Ошибка на 238 строке:")
			fmt.Println(err.Error())
		}
		_, err = d.Exec("insert into messages values ($1, $2, $3)", ms.User, ms.Chat, ms.Data)
		if err != nil {
			fmt.Print("Ошибка на 243 строке:")
			fmt.Println(err.Error())
			return
		}
		err = d.Close()
		if err != nil {
			fmt.Print("Ошибка на 249 строке:")
			fmt.Println(err.Error())
		}
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			var (
				chatList []string
			)
			db, err := sql.Open("sqlite3", "data/intensive.db")
			defer db.Close()
			if err != nil {
				fmt.Print("Ошибка на 258 строке:")
				fmt.Println(err.Error())
				return false
			}
			r, err := db.Query("select name from chat")
			if err != nil {
				fmt.Print("Ошибка на 265 строке:")
				fmt.Println(err.Error())
				return false
			}
			for r.Next() {
				var name string
				err = r.Scan(&name)
				if err != nil {
					fmt.Println(err.Error())
					return false
				}
				chatList = append(chatList, name)
			}
			url := strings.Split(s.Request.URL.Path, "/")
			chatNameS := url[len(url)-2]
			url = strings.Split(q.Request.URL.Path, "/")
			chatNameQ := url[len(url)-2]
			allowedInDb := slices.IndexFunc(chatList, func(e string) bool {
				return chatNameQ == e
			})
			if allowedInDb != -1 {
				return chatNameS == chatNameQ
			}
			return false
		})
	})
}
