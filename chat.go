package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"gopkg.in/olahol/melody.v1"
	"net/http"
	"strconv"
	"strings"

	. "intensive.com/data"
	. "intensive.com/errors"
	. "intensive.com/models"
)

type DTO struct {
	name string `json:"name" binding:"require"`
}

var m = melody.New()

func ConfigureChatControllers(r *gin.Engine) {

	r.GET("/chat/:id", GetChat)

	r.GET("/id/:mod/:name", GetObjectId)

	r.GET("/chat/:id/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	r.GET("/channel/massages/:chatId/:period", GetChatMessages)

	m.HandleMessage(HandleMessages)
}

func GetChat(c *gin.Context) {
	var (
		chatList []int
	)
	r, err := Db.Query("select id from chat")
	if err != nil {
		fmt.Println("chat:49")
		go GlobalErrorsHandler.SendError(Error{
			Err:      err,
			Location: "chat.go >> line 44",
		})
		c.HTML(505, "mistake.html", nil)
	}
	for r.Next() {
		var name int

		if err = r.Scan(&name); err != nil {
			fmt.Println("chat:49")
			go GlobalErrorsHandler.SendError(Error{
				Err:      err,
				Location: "chat.go >> line 56",
			})
			c.HTML(505, "mistake.html", nil)
		}
		chatList = append(chatList, name)
	}
	id := c.Param("id")
	name, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("chat:49")
		go GlobalErrorsHandler.SendError(Error{
			Err:      err,
			Location: "chat.go >> line 67",
		})
		c.HTML(505, "mistake.html", nil)
	}
	allowedInDb := slices.IndexFunc(chatList, func(e int) bool {
		return name == e
	})
	if allowedInDb != -1 {
		c.HTML(http.StatusOK, "chat.html", nil)
		return
	} else {
		c.HTML(505, "mistake.html", nil)
		go GlobalErrorsHandler.SendError(Error{
			Err:      "Chat name not allowed in DataBase",
			Location: "chat.go >> line 44",
		})
		return
	}
}

func GetObjectId(c *gin.Context) {
	if c.Param("mod") == "user" {
		var dto DTO
		var uid int
		dto.name = c.Param("name")
		r, err := Db.Query("select id from users where name = $1", dto.name)
		if err != nil {
			fmt.Print("Ошибка на 97 строке:")
			go GlobalErrorsHandler.SendError(Error{
				Err:      err,
				Location: "chat.go >> line 97",
			})
			fmt.Println(err.Error())
		}
		for r.Next() {
			if err = r.Scan(&uid); err != nil {
				fmt.Print("Ошибка на 103 строке:")
				go GlobalErrorsHandler.SendError(Error{
					Err:      err,
					Location: "chat.go >> line 107",
				})
				fmt.Println(err.Error())
			}
		}
		c.JSON(200, map[string]any{
			"id": uid,
		})
	} else if c.Param("mod") == "chat" {
		var dto DTO
		var cid int
		dto.name = c.Param("name")
		r, err := Db.Query("select id from chat where name = $1", dto.name)
		if err != nil {
			fmt.Print("Ошибка на 126 строке:")
			go GlobalErrorsHandler.SendError(Error{
				Err:      err,
				Location: "chat.go >> line 123",
			})
			fmt.Println(err.Error())
		}
		for r.Next() {
			if err = r.Scan(&cid); err != nil {
				fmt.Print("Ошибка на 132 строке:")
				go GlobalErrorsHandler.SendError(Error{
					Err:      err,
					Location: "chat.go >> line 133",
				})
				fmt.Println(err.Error())
			}
			fmt.Println(cid)
		}
		c.JSON(200, map[string]any{
			"id": cid,
		})
	} else {

	}
}

func GetChatMessages(c *gin.Context) {
	response := make(map[string]MsgDTO)
	r, err := Db.Query("select fr, message from messages where ch = $1", c.Param("chatId"))
	if err != nil {
		fmt.Print("Ошибка на 163 строке:")
		go GlobalErrorsHandler.SendError(Error{
			Err:      err,
			Location: "chat.go >> line 153",
		})
		fmt.Println(err.Error())
	}
	i := 1
	for r.Next() {
		var (
			msg    MsgDTO
			userId int
		)
		if err = r.Scan(&userId, &msg.Data); err != nil {
			fmt.Print("Ошибка на 174 строке:")
			go GlobalErrorsHandler.SendError(Error{
				Err:      err,
				Location: "chat.go >> line 168",
			})
			fmt.Println(err.Error())
		}
		re, err := Db.Query("select name from users where id = $1", userId)
		if err != nil {
			fmt.Print("Ошибка на 179 строке:")
			go GlobalErrorsHandler.SendError(Error{
				Err:      err,
				Location: "chat.go >> line 176",
			})
			fmt.Println(err.Error())
		}
		for re.Next() {
			if err = re.Scan(&msg.User); err != nil {
				go GlobalErrorsHandler.SendError(Error{
					Err:      err,
					Location: "chat.go >> line 186",
				})
			}
		}
		response[string(i)] = msg
		fmt.Println(response[string(i)])
		i = i + 1
	}
	c.JSON(200, response)
}

func HandleMessages(s *melody.Session, msg []byte) {
	var dto MsgDTO
	var ms Msg
	_ = json.Unmarshal(msg, &dto)
	ms.Chat, _ = strconv.ParseInt(dto.Chat, 10, 0)
	ms.User, _ = strconv.ParseInt(dto.User, 10, 0)
	ms.Data = dto.Data
	fmt.Println(ms)

	r, err := Db.Query("select name from chat where id = $1", ms.Chat)
	if err != nil {
		fmt.Print("Ошибка на 214 строке:")
		go GlobalErrorsHandler.SendError(Error{
			Err:      err,
			Location: "chat.go >> line 209",
		})
		fmt.Println(err.Error())
		return
	}

	re, err := Db.Query("select name from users where id = $1", ms.User)
	if err != nil {
		fmt.Print("Ошибка на 220 строке:")
		go GlobalErrorsHandler.SendError(Error{
			Err:      err,
			Location: "chat.go >> line 220",
		})
		fmt.Println(err.Error())
		return
	}

	for r.Next() && re.Next() {
		if err = r.Scan(&dto.Chat); err != nil {
			fmt.Print("Ошибка на 227 строке:")
			go GlobalErrorsHandler.SendError(Error{
				Err:      err,
				Location: "chat.go >> line 232",
			})
			fmt.Println(err.Error())
		}
		if err = re.Scan(&dto.User); err != nil {
			fmt.Print("Ошибка на 232 строке:")
			go GlobalErrorsHandler.SendError(Error{
				Err:      err,
				Location: "chat.go >> line 240",
			})
			fmt.Println(err.Error())
		}
	}
	if err = r.Close(); err != nil {
		fmt.Print("Ошибка на 238 строке:")
		go GlobalErrorsHandler.SendError(Error{
			Err:      err,
			Location: "chat.go >> line 249",
		})
		fmt.Println(err.Error())
	}
	if err = re.Close(); err != nil {
		fmt.Print("Ошибка на 238 строке:")
		go GlobalErrorsHandler.SendError(Error{
			Err:      err,
			Location: "chat.go >> line 257",
		})
		fmt.Println(err.Error())
	}
	msg, err = json.Marshal(dto)
	if err != nil {
		fmt.Print("Ошибка на 238 строке:")
		go GlobalErrorsHandler.SendError(Error{
			Err:      err,
			Location: "chat.go >> line 265",
		})
		fmt.Println(err.Error())
	}
	if _, err = Db.Exec("insert into messages values ($1, $2, $3)", ms.User, ms.Chat, ms.Data); err != nil {
		fmt.Print("Ошибка на 238 строке:")
		go GlobalErrorsHandler.SendError(Error{
			Err:      err,
			Location: "chat.go >> line 274",
		})
		fmt.Println(err.Error())
	}

	m.BroadcastFilter(msg, func(q *melody.Session) bool {
		var chatList []int

		chatList, err = BroadcastDbWork()
		if err != nil {
			return false
		}

		url := strings.Split(s.Request.URL.Path, "/")
		chatNameS, err := strconv.Atoi(url[len(url)-2])
		if err != nil {
			fmt.Print("Ошибка на 265 строке:")
			go GlobalErrorsHandler.SendError(Error{
				Err:      err,
				Location: "chat.go >> line 298",
			})
			fmt.Println(err.Error())
			return false
		}

		url = strings.Split(q.Request.URL.Path, "/")
		chatNameQ, err := strconv.Atoi(url[len(url)-2])
		if err != nil {
			fmt.Print("Ошибка на 265 строке:")
			go GlobalErrorsHandler.SendError(Error{
				Err:      err,
				Location: "chat.go >> line 310",
			})
			fmt.Println(err.Error())
			return false
		}

		allowedInDb := slices.IndexFunc(chatList, func(e int) bool {
			return chatNameQ == e
		})
		if allowedInDb != -1 {
			return chatNameS == chatNameQ
		}

		go GlobalErrorsHandler.SendError(Error{
			Err:      "Chat Name not allowed in DataBase",
			Location: "chat.go >> line 331",
		})

		return false
	})
}

func BroadcastDbWork() ([]int, error) {
	var chatList []int
	r, err := Db.Query("select id from chat")
	if err != nil {
		fmt.Print("Ошибка на 265 строке:")
		go GlobalErrorsHandler.SendError(Error{
			Err:      err,
			Location: "chat.go >> line 340",
		})
		fmt.Println(err.Error())
		return nil, err
	}
	for r.Next() {
		var name int
		err = r.Scan(&name)
		if err != nil {
			fmt.Println(err.Error())
			go GlobalErrorsHandler.SendError(Error{
				Err:      err,
				Location: "chat.go >> line 351",
			})
			return nil, err
		}
		chatList = append(chatList, name)
	}
	return chatList, nil
}
