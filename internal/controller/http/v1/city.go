package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"lg/internal/usecase"
	"net/http"
)

type cityRoutes struct {
	c usecase.CityContract
}

type cityDTO struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type citiListResponse struct {
	Cities []cityDTO `json:"cities"`
}

func NewCityRoutes(handler *gin.RouterGroup, c usecase.CityContract) {
	cr := &cityRoutes{c: c}
	handler.GET("/cities/:uuid", cr.getCityByCountryUUID)
}

func (cr *cityRoutes) getCityByCountryUUID(c *gin.Context) {
	countryKey, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	cities, err := cr.c.GetCitiesByCountryUUID(c.Request.Context(), countryKey)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var responseList []cityDTO
	for _, v := range cities {
		responseList = append(responseList, cityToDTO(v))
	}
	c.JSON(http.StatusOK, citiListResponse{responseList})
}
