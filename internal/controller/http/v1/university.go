package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
	"net/http"
)

type universityRoutes struct {
	c usecase.UniversityContract
}

type universityDTO struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type universityListResponse struct {
	Universities []universityDTO `json:"universities"`
}

func newUniversitiesRoutes(handler *gin.RouterGroup, c usecase.UniversityContract) {
	cr := &universityRoutes{c: c}
	handler.GET("/university", cr.getAllUniversities)
}

func (cr *universityRoutes) getAllUniversities(c *gin.Context) {
	universityList, err := cr.c.GetAllUniversities(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var responseList []universityDTO
	for _, v := range universityList {
		responseList = append(responseList, universityToDTO(v))
	}
	c.JSON(http.StatusOK, universityListResponse{responseList})
}
