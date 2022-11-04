package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
	"net/http"
)

type countryRoutes struct {
	c usecase.CountryContract
}

type countryDTO struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type countryListResponse struct {
	Countries []countryDTO `json:"countries"`
}

func newCountryRoute(handler *gin.RouterGroup, c usecase.CountryContract) {
	cr := &countryRoutes{c: c}
	handler.GET("/country", cr.getAllCountries)
}

func (cr *countryRoutes) getAllCountries(c *gin.Context) {
	countryList, err := cr.c.GetAllCountries(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var responseList []countryDTO
	for _, v := range countryList {
		responseList = append(responseList, countryToDTO(v))
	}
	c.JSON(http.StatusOK, countryListResponse{responseList})
}
