package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
	"net/http"
)

type eduspecialityRoutes struct {
	c usecase.EduspecialityContract
}

type eduspecialityDTO struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type eduspecialityListResponse struct {
	Eduspecialities []eduspecialityDTO `json:"eduspecialities"`
}

func newEduspecialitiesRoutes(handler *gin.RouterGroup, c usecase.EduspecialityContract) {
	cr := &eduspecialityRoutes{c: c}
	handler.GET("/eduspeciality", cr.getAllEduspecialities)
}

func (cr *eduspecialityRoutes) getAllEduspecialities(c *gin.Context) {
	eduspecialityList, err := cr.c.GetAllEduspecialities(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var responseList []eduspecialityDTO
	for _, v := range eduspecialityList {
		responseList = append(responseList, eduspecialityToDTO())
	}
	c.JSON(http.StatusOK, eduspecialityListResponse{responseList})
}
