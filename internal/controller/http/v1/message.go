package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"lg/internal/entity"
	"lg/internal/usecase"
	"net/http"
	"time"
)

type messageRoutes struct {
	j  usecase.JwtContract
	mg usecase.MessageContract
	ch usecase.ChatContract
}

func newMessageRoutes(handler *gin.RouterGroup, j usecase.JwtContract, mg usecase.MessageContract, ch usecase.ChatContract) {
	m := messageRoutes{j: j, mg: mg, ch: ch}

	handler.POST("/send-message", m.storeMessage)
	handler.POST("/rollup", m.rollUp)
}

type messageRequest struct {
	Content  string    `json:"content"`
	Type     string    `json:"type"`
	ChatUUID uuid.UUID `json:"chatUUID"`
}

func (m *messageRoutes) storeMessage(c *gin.Context) {
	access, err := c.Cookie("access")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	tokenUUID, err := m.j.CheckToken(access)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userUUID, err := uuid.Parse(tokenUUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var msgRequest messageRequest
	if err := c.ShouldBindJSON(&msgRequest); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = m.mg.StoreMessage(c.Request.Context(), entity.Message{
		AuthorUUID:   userUUID,
		Type:         msgRequest.Type,
		Content:      msgRequest.Content,
		CreationDate: time.Now(),
		ChatUUID:     msgRequest.ChatUUID,
	})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

type adminMessageRequest struct {
	CreatorUUID string `json:"creatorUUID"` // создатель проекта
}

// создание чата с админом
func (m *messageRoutes) rollUp(c *gin.Context) {
	access, err := c.Cookie("access")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	tokenUUID, err := m.j.CheckToken(access)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userUUID, err := uuid.Parse(tokenUUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var adm adminMessageRequest
	if err := c.ShouldBindJSON(&adm); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	creatorUUID, err := uuid.Parse(adm.CreatorUUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	chatUUID, err := m.ch.CreateChat(c.Request.Context(), "", []uuid.UUID{userUUID, creatorUUID})
	err = m.mg.StoreMessage(c.Request.Context(), entity.Message{
		AuthorUUID:   userUUID,
		Type:         "invite",
		Content:      "Прив, го вирт по вебке?",
		CreationDate: time.Now(),
		ChatUUID:     chatUUID,
	})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
