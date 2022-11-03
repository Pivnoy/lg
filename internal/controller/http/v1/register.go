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

	handler.POST("/register", s.register)
}

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerResponse struct {
	UUID string `json:"uuid"`
}

func (s *registerContract) register(c *gin.Context) {
	var request registerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	ud, err := s.s.CreateNewUser(c.Request.Context(), request.Email, request.Password)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, registerResponse{ud.String()})
}
