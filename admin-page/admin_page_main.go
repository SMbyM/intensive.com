package admin_page

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	. "intensive.com/errors"
)

type MessageType int16

var (
	getMessages MessageType = 1
	getUsers    MessageType = 2
	setUsers    MessageType = 3
	modifyUsers MessageType = 4

	errorMessage MessageType = 5
	newMessage   MessageType = 6
)

type Data struct {
	Content []byte `json:"content"`
}

type Message struct {
	Message string `json:"message"`

	Type MessageType `json:"type"`

	Data Data `json:"data"`
}

func ConfigureAdminPage(r *gin.Engine) {
	m := melody.New()

	r.GET("/admin/admin_page", GetAdminPage)
	r.POST("/admin/login_admin", LoginAdmin)
	r.GET("/admin/stats", GetAdminStats)
	r.GET("admin/messages/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		var msg_ Message
		if err := json.Unmarshal(msg, &msg_); err != nil {
			go GlobalErrorsHandler.SendError(Error{
				Err:      err,
				Location: "package admin-page >> admin_page_main.go >> line 59",
			})
		}

		switch msg_.Type {
		case getUsers:
			m.Broadcast(AdminGetUsers())
		case getMessages:
			m.Broadcast(AdminGetMessages())
		case setUsers:
			m.Broadcast(AdminSetUsers())
		case modifyUsers:
			m.Broadcast(AdminModifyUsers())
		}
	})

	go func() {
		for {
			chanErr := <-GlobalErrorsHandler.Errors
			sentErr, err := json.Marshal(&chanErr)
			if err != nil {
				go GlobalErrorsHandler.SendError(Error{
					Err:      err,
					Location: "package admin-page >> admin_page_main.go >> line 70",
				})
			}
			m.Broadcast(sentErr)
		}
	}()

}

func AdminModifyUsers() []byte {
	return []byte{}
}

func AdminSetUsers() []byte {
	return []byte{}
}

func AdminGetMessages() []byte {
	return []byte{}
}

func AdminGetUsers() []byte {
	return []byte{}
}

func GetAdminPage(c *gin.Context) {

}

func LoginAdmin(c *gin.Context) {

}

func GetAdminStats(c *gin.Context) {

}
