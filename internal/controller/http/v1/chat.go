package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
	"net/http"
)

type chatRoutes struct {
	c usecase.ChatContract
	j usecase.JwtContract
}

//type ChatItem = {
//chat_name: string
//chat_uuid: string
//lastMessage: Message
//imageUrl?: string // только в групповом чате - ссылка на картинку проекта
//}
//
//type Chat = {
//uuid: string
//projectUuid?: string // только в групповом чате
//history?: ChatHistory
//type: 'group' | 'direct'
//users: User[]
//}
//
//content: string
//sender: User
//date: string

func newChatRoutes(handler *gin.RouterGroup, c usecase.ChatContract) {
	ch := chatRoutes{c: c}

	handler.POST("/getChats", ch.getChatList)
}

func (ch *chatRoutes) getChatList(c *gin.Context) {
	access, err := c.Cookie("access")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// Add here token with UUID
	_, err = ch.j.CheckToken(access)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

}
