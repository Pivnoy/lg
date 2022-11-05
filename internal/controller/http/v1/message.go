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
}

func newMessageRoutes(handler *gin.RouterGroup, j usecase.JwtContract, mg usecase.MessageContract) {
	m := messageRoutes{j: j, mg: mg}

	handler.POST("/send-message", m.storeMessage)
}

type messageRequest struct {
	Content  string    `json:"content"`
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
