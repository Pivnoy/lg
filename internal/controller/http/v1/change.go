package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"lg/internal/entity"
	"lg/internal/usecase"
	"net/http"
)

type changeRoutes struct {
	j usecase.JwtContract
	u usecase.UserContract
}

func newChangeRoutes(handler *gin.RouterGroup, j usecase.JwtContract, u usecase.UserContract) {
	c := changeRoutes{j: j, u: u}

	handler.POST("/change-password", c.changePassword)
}

type changeRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func (ch *changeRoutes) changePassword(c *gin.Context) {
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
	var change changeRequest
	if err := c.ShouldBindJSON(&change); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	us, err := ch.u.GetUserByUUID(c.Request.Context(), userUUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = ch.j.CompareUserPassword(c.Request.Context(), entity.User{
		Email:    us.Email,
		Password: change.OldPassword,
	})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = ch.u.ChangePassword(c.Request.Context(), us.Email, change.NewPassword)
	c.JSON(http.StatusOK, nil)
}
