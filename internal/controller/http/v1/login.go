package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/entity"
	"lg/internal/usecase"
	"net/http"
)

type loginRoutes struct {
	j usecase.JwtContract
}

func newLoginRoutes(handler *gin.RouterGroup, j usecase.JwtContract) {
	r := &loginRoutes{j: j}

	handler.POST("/login", r.login)
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *loginRoutes) login(c *gin.Context) {
	var lReq loginRequest
	if err := c.ShouldBindJSON(&lReq); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := l.j.CompareUserPassword(c.Request.Context(), entity.User{
		Email:    lReq.Email,
		Password: lReq.Password,
	})
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Cannot find user in db or cmp psswd")
	}
	token, err := l.j.GenerateToken(lReq.Email)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.SetCookie("access", token, int(120*60), "/", "", false, false)
	c.JSON(http.StatusOK, token)
}
