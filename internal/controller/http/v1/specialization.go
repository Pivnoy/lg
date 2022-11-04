package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
	"net/http"
)

type specializationRoutes struct {
	c usecase.SpecializationContract
}

type specializationDTO struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type specializationListResponse struct {
	Specializations []specializationDTO `json:"specializations"`
}

func newSpecializationsRoutes(handler *gin.RouterGroup, c usecase.SpecializationContract) {
	cr := &specializationRoutes{c: c}
	handler.GET("/specialization", cr.getAllSpecializations)
}
func (cr *specializationRoutes) getAllSpecializations(c *gin.Context) {
	specializationList, err := cr.c.GetAllSpecializations(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var responseList []specializationDTO
	for _, v := range specializationList {
		responseList = append(responseList, specializationToDTO(v))
	}
	c.JSON(http.StatusOK, specializationListResponse{responseList})
}
