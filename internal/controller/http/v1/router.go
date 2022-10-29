package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
)

//signin
//login
//logout
//profile

func NewRouter(handler *gin.Engine,
	p usecase.ProjectContract,
) {
	h := handler.Group("/api/v1")
	{
		newProjectRouter(h, p)
	}
}
