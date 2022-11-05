package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
	"net/http"
)

type companyRoutes struct {
	cs usecase.CompanyContract
}

type reqInnDTO struct {
	Inn string `json:"inn"`
}

type resInnDTO struct {
	Result bool `json:"result"`
}

func newCompanyRoutes(handler *gin.RouterGroup, cs usecase.CompanyContract) {
	cr := &companyRoutes{cs: cs}
	handler.POST("/company/inn", cr.checkInn)
}

func (cr *companyRoutes) checkInn(c *gin.Context) {
	req := new(reqInnDTO)
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	res, err := cr.cs.CheckCompanyExistenceByInn(c.Request.Context(), req.Inn)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, resInnDTO{res})
}
