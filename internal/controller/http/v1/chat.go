package v1

import (
	"fmt"
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
	handler.POST("/kick", ch.kickUser)
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

type kickRequest struct {
	UserUUID string `json:"userUUID"`
	ChatUUID string `json:"chatUUID"`
}

func (ch *chatRoutes) kickUser(c *gin.Context) {
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
	userUUID, err := uuid.Parse(tokenUUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var kick kickRequest
	if err := c.ShouldBindJSON(&kick); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	chatUUID, err := uuid.Parse(kick.ChatUUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	dropUserUUID, err := uuid.Parse(kick.UserUUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	creatorUUID, err := ch.c.GetCreatorByChat(c.Request.Context(), chatUUID)
	fmt.Println(creatorUUID)
	fmt.Println(dropUserUUID)
	fmt.Println(userUUID)
	if creatorUUID == userUUID && creatorUUID != dropUserUUID {
		err = ch.c.DeleteUserFromChat(c.Request.Context(), chatUUID, dropUserUUID)
		if err != nil {
			errorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	} else if userUUID == dropUserUUID {
		err = ch.c.DeleteUserFromChat(c.Request.Context(), chatUUID, dropUserUUID)
		if err != nil {
			errorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, nil)
}
