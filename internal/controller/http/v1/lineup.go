package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"lg/internal/usecase"
	"net/http"
)

type lineupRoutes struct {
	uc usecase.LineupContract
	j  usecase.JwtContract
}

func newLineupRoutes(handler *gin.RouterGroup, uc usecase.LineupContract, j usecase.JwtContract) {
	lr := &lineupRoutes{uc: uc, j: j}
	handler.PUT("/lineup/:uuid", lr.addCommand)
}

func (lr *lineupRoutes) addCommand(c *gin.Context) {
	access, err := c.Cookie("access")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	tokenUUID, err := lr.j.CheckToken(access)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userUUID, err := uuid.Parse(tokenUUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	lineupKey, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = lr.uc.UpdateLineup(c.Request.Context(), lineupKey, userUUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
