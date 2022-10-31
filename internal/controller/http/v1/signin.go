package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
	"net/http"
)

type signInRoutes struct {
	s usecase.SignInContract
}

func newSignInRoutes(handler *gin.RouterGroup, si usecase.SignInContract) {
	s := signInRoutes{s: si}

	handler.POST("/signin", s.SignIn)
}

type signInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *signInRoutes) SignIn(c *gin.Context) {
	var request signInRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := s.s.CreateNewUser(c.Request.Context(), request.Username, request.Password)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
