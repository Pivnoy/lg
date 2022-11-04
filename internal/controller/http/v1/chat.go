package v1

import (
	"github.com/gin-gonic/gin"
)

type chatRoutes struct {
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

func newChatRoutes(handler *gin.RouterGroup) {

}
