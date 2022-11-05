package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
	"net/http"
)

type categoryRoutes struct {
	cu usecase.CategoryContract
	j  usecase.JwtContract
}

type categoryDTO struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type categoryListResponse struct {
	Categories []categoryDTO `json:"categories"`
}

func newCategoryRoutes(handler *gin.RouterGroup, cu usecase.CategoryContract, j usecase.JwtContract) {
	cr := &categoryRoutes{cu: cu, j: j}
	handler.GET("/category", cr.getAllCategory)
}

// @Summary GetAllCategory
// @Tags Categories
// @Description Get all categories
// @Success 200 {object} categoryListResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/category [get]
func (cr *categoryRoutes) getAllCategory(c *gin.Context) {
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
	categoryList, err := cr.cu.GetAllCategory(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var responseList []categoryDTO
	for _, v := range categoryList {
		responseList = append(responseList, categoryToDTO(v))
	}
	c.JSON(http.StatusOK, categoryListResponse{responseList})
}
