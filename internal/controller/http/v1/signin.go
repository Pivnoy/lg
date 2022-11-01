package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
	"net/http"
)

type registerContract struct {
	s usecase.RegisterContract
}

func newRegisterRoutes(handler *gin.RouterGroup, si usecase.RegisterContract) {
	s := registerContract{s: si}

	handler.POST("/register", s.SignIn)
}

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *registerContract) SignIn(c *gin.Context) {
	var request registerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := s.s.CreateNewUser(c.Request.Context(), request.Email, request.Password)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
