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
}

func newChatRoutes(handler *gin.RouterGroup, c usecase.ChatContract, j usecase.JwtContract) {
	ch := chatRoutes{c: c, j: j}

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
		chatsItems = append(chatsItems, chatItemToDTO(chat))
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
		msgDTO = append(msgDTO, messageToDTO(msg))
	}
	c.JSON(http.StatusOK, chatHistoryDTO{msgDTO})
}
