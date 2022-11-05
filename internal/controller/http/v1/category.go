package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
	"net/http"
)

type categoryRoutes struct {
	cu usecase.CategoryContract
}

type categoryDTO struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type categoryListResponse struct {
	Categories []categoryDTO `json:"categories"`
}

func newCategoryRoutes(handler *gin.RouterGroup, cu usecase.CategoryContract) {
	cr := &categoryRoutes{cu: cu}
	handler.GET("/category", cr.getAllCategory)
}

func (cr *categoryRoutes) getAllCategory(c *gin.Context) {
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
