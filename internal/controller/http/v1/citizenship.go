package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
	"net/http"
)

type citizenshipRoutes struct {
	c usecase.CitizenshipContract
}

type citizenshipDTO struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type citizenshipListResponse struct {
	Citizenships []citizenshipDTO `json:"citizenships"`
}

func newCitizenshipRoutes(handler *gin.RouterGroup, c usecase.CitizenshipContract) {
	cr := &citizenshipRoutes{c: c}
	handler.GET("/citizenship", cr.getAllCitizenships)
}

func (cr *citizenshipRoutes) getAllCitizenships(c *gin.Context) {
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
