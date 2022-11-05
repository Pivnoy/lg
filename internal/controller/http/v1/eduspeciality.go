package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
	"net/http"
)

type eduspecialityRoutes struct {
	c usecase.EduspecialityContract
	j usecase.JwtContract
}

type eduspecialityDTO struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type eduspecialityListResponse struct {
	Eduspecialities []eduspecialityDTO `json:"eduspecialities"`
}

func newEduspecialitiesRoutes(handler *gin.RouterGroup, c usecase.EduspecialityContract, j usecase.JwtContract) {
	cr := &eduspecialityRoutes{c: c, j: j}
	handler.GET("/eduspeciality", cr.getAllEduspecialities)
}

func (cr *eduspecialityRoutes) getAllEduspecialities(c *gin.Context) {
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
	eduspecialityList, err := cr.c.GetAllEduspecialities(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var responseList []eduspecialityDTO
	for _, v := range eduspecialityList {
		responseList = append(responseList, eduspecialityToDTO(v))
	}
	c.JSON(http.StatusOK, eduspecialityListResponse{responseList})
}
