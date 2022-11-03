package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
)

func NewRouter(handler *gin.Engine,
	p usecase.ProjectContract,
	s usecase.RegisterContract,
	j usecase.JwtContract,
	u usecase.UserContract,
	pr usecase.ProfileContract,
) {
	h := handler.Group("/api/v1")
	{
		newProjectRouter(h, p)
		newRegisterRoutes(h, s)
		newLoginRoutes(h, j, u, pr)
		newLogoutRouter(h)
	}
}
