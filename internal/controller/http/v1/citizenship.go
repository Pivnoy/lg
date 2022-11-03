package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"lg/internal/usecase"
	"net/http"
)

type citizenshipRoute struct {
	c usecase.CitizenshipContract
}

type citizenshipDTO struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	CountryUUID uuid.UUID `json:"country_uuid"`
}

type citizenshipListResponse struct {
	Citizenships []citizenshipDTO `json:"citizenships"`
}

func newCitizenshipRoute(handler *gin.RouterGroup, c usecase.CitizenshipContract) {
	cr := &citizenshipRoute{c: c}
	handler.GET("/citizenship", cr.getAllCitizenships)
}

func (cr *citizenshipRoute) getAllCitizenships(c *gin.Context) {
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
