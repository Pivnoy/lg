package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
	"net/http"
)

type citizenshipRoutes struct {
	c usecase.CitizenshipContract
	j usecase.JwtContract
}

type citizenshipDTO struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type citizenshipListResponse struct {
	Citizenships []citizenshipDTO `json:"citizenships"`
}

func newCitizenshipRoutes(handler *gin.RouterGroup, c usecase.CitizenshipContract, j usecase.JwtContract) {
	cr := &citizenshipRoutes{c: c, j: j}
	handler.GET("/citizenship", cr.getAllCitizenships)
}

func (cr *citizenshipRoutes) getAllCitizenships(c *gin.Context) {
	access, err := c.Cookie("access")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	_, err = cr.j.CheckToken(access)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	citizenshipList, err := cr.c.GetAllCitizenships(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var responseList []citizenshipDTO
	for _, v := range citizenshipList {
		responseList = append(responseList, citizenshipToDTO(v))
	}
	c.JSON(http.StatusOK, citizenshipListResponse{responseList})
}
