package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
	"net/http"
)

type employmentRoutes struct {
	c usecase.EmploymentContract
	j usecase.JwtContract
}

type employmentDTO struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type employmentListResponse struct {
	Employments []employmentDTO `json:"employments"`
}

func newEmploymentsRoutes(handler *gin.RouterGroup, c usecase.EmploymentContract, j usecase.JwtContract) {
	cr := &employmentRoutes{c: c, j: j}
	handler.GET("/employment", cr.getAllEmployments)
}

func (cr *employmentRoutes) getAllEmployments(c *gin.Context) {
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
	employmentList, err := cr.c.GetAllEmployments(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var responseList []employmentDTO
	for _, v := range employmentList {
		responseList = append(responseList, employmentToDTO(v))
	}
	c.JSON(http.StatusOK, employmentListResponse{responseList})
}
