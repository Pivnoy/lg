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

//TODO Patronymic
type loginResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UUID      string `json:"uuid"`
}

//TODO user doesn't exists
func (l *loginRoutes) login(c *gin.Context) {
	var lReq loginRequest
	if err := c.ShouldBindJSON(&lReq); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err := l.j.CompareUserPassword(c.Request.Context(), entity.User{
		Email:    lReq.Email,
		Password: lReq.Password,
	})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "Cannot find user in db or cmp psswd")
	}
	token, err := l.j.GenerateToken(lReq.Email)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.SetCookie("access", token, 120*60, "/", "", false, true)
	c.JSON(http.StatusOK, loginResponse{})
}
