package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"lg/internal/usecase"
	"net/http"
)

type countryRoute struct {
	c usecase.CountryContract
}

type countryDTO struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
	Code string    `json:"code"`
}

type countryListResponse struct {
	Countries []countryDTO `json:"countries"`
}

func newCountryRoute(handler *gin.RouterGroup, c usecase.CountryContract) {
	cr := &countryRoute{c: c}
	handler.GET("/country", cr.getAllCountries)
}

func (cr *countryRoute) getAllCountries(c *gin.Context) {
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
