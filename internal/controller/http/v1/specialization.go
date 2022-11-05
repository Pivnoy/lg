package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
	"net/http"
)

type specializationRoutes struct {
	c usecase.SpecializationContract
	j usecase.JwtContract
}

type specializationDTO struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type specializationListResponse struct {
	Specializations []specializationDTO `json:"specializations"`
}

func newSpecializationsRoutes(handler *gin.RouterGroup, c usecase.SpecializationContract, j usecase.JwtContract) {
	cr := &specializationRoutes{c: c, j: j}
	handler.GET("/specialization", cr.getAllSpecializations)
}

// @Summary GetAllSpecializations
// @Tags Specializations
// @Description Get all specializations
// @Success 200 {object} specializationListResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/specialization [get]
func (cr *specializationRoutes) getAllSpecializations(c *gin.Context) {
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
