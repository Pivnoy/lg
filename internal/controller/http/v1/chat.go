package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"lg/internal/usecase"
	"net/http"
)

type chatRoutes struct {
	c usecase.ChatContract
	j usecase.JwtContract
	p usecase.ProfileContract
	m usecase.MessageContract
}

func newChatRoutes(handler *gin.RouterGroup, c usecase.ChatContract, j usecase.JwtContract, p usecase.ProfileContract, m usecase.MessageContract) {
	ch := chatRoutes{c: c, j: j, p: p, m: m}

	handler.POST("/chats", ch.getChatList)
	handler.POST("/history", ch.getChatHistory)
}

type chatItemResponse struct {
	ChatItems []chatItemDTO `json:"chatItems"`
}

func (ch *chatRoutes) getChatList(c *gin.Context) {
	access, err := c.Cookie("access")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	tokenUUID, err := ch.j.CheckToken(access)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	user, err := uuid.Parse(tokenUUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	chats, err := ch.c.GetAllChatsByUser(c.Request.Context(), user)
	var chatsItems []chatItemDTO
	for _, chat := range chats {
		lastMsg, err := ch.m.GetLastMessageByChat(c.Request.Context(), chat.ChatUUID)
		if err != nil {
			errorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		prf, err := ch.p.GetProfileByUser(c.Request.Context(), lastMsg.AuthorUUID)
		if err != nil {
			errorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		chatsItems = append(chatsItems, chatItemToDTO(chat, lastMsg, userToDTO(prf)))
	}
	c.JSON(http.StatusOK, chatItemResponse{chatsItems})
}

type chatHistoryRequest struct {
	ChatUUID string `json:"chatUUID"`
}

func (ch *chatRoutes) getChatHistory(c *gin.Context) {
	access, err := c.Cookie("access")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	_, err = ch.j.CheckToken(access)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var chatHst chatHistoryRequest
	if err := c.ShouldBindJSON(&chatHst); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	chatUUID, err := uuid.Parse(chatHst.ChatUUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	history, err := ch.c.GetChatHistory(c.Request.Context(), chatUUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var msgDTO []messageDTO
	for _, msg := range history {
		prf, err := ch.p.GetProfileByUser(c.Request.Context(), msg.AuthorUUID)
		us := userToDTO(prf)
		if err != nil {
			errorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		msgDTO = append(msgDTO, messageToDTO(msg, us))
	}
	c.JSON(http.StatusOK, chatHistoryDTO{msgDTO})
}
