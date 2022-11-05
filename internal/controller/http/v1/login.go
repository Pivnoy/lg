package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lg/internal/entity"
	"lg/internal/usecase"
	"net/http"
)

type loginRoutes struct {
	j usecase.JwtContract
	u usecase.UserContract
	p usecase.ProfileContract
}

func newLoginRoutes(handler *gin.RouterGroup, j usecase.JwtContract, u usecase.UserContract, p usecase.ProfileContract) {
	r := &loginRoutes{j: j, u: u, p: p}

	handler.POST("/login", r.login)
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	UUID       string `json:"uuid"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Patronymic string `json:"patronymic"`
}

func (l *loginRoutes) login(c *gin.Context) {
	var lReq loginRequest
	if err := c.ShouldBindJSON(&lReq); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	us, err := l.u.GetUser(c.Request.Context(), lReq.Email)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(us)
	err = l.j.CompareUserPassword(c.Request.Context(), entity.User{
		Email:    lReq.Email,
		Password: lReq.Password,
	})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "Cannot find user in db or cmp psswd")
		return
	}
	token, err := l.j.GenerateToken(us.UUID.String())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	prf, err := l.p.GetProfileByUser(c.Request.Context(), us.UUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.SetCookie("access", token, 120*60, "/", "", false, true)
	c.JSON(http.StatusOK, loginResponse{
		UUID:       prf.UserUUID.String(),
		FirstName:  prf.Firstname,
		LastName:   prf.Lastname,
		Patronymic: prf.Patronymic.String,
	})
}
