package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type logoutRouter struct {
}

func newLogoutRouter(handler *gin.RouterGroup) {
	l := logoutRouter{}

	handler.POST("/logout", l.logout)
}

func (l *logoutRouter) logout(c *gin.Context) {
	c.SetCookie("access", "", -1, "/", "", false, false)
	c.JSON(http.StatusOK, nil)
}
