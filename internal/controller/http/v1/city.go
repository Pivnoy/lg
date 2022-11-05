package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"lg/internal/usecase"
	"net/http"
)

type cityRoutes struct {
	c usecase.CityContract
	j usecase.JwtContract
}

type cityDTO struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type citiListResponse struct {
	Cities []cityDTO `json:"cities"`
}

func newCityRoutes(handler *gin.RouterGroup, c usecase.CityContract, j usecase.JwtContract) {
	cr := &cityRoutes{c: c, j: j}
	handler.GET("/cities/:uuid", cr.getCityByCountryUUID)
}

// @Summary GetCityByCountry
// @Tags Cities
// @Description Get cities by country
// @Param uuid path string true "Enter uuid country"
// @Success 200 {object} citiListResponse
// @Failure 400 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/cities/{uuid} [get]
func (cr *cityRoutes) getCityByCountryUUID(c *gin.Context) {
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
