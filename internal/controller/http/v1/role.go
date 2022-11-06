package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
	"net/http"
)

type roleRoutes struct {
	ru usecase.RoleContract
	j  usecase.JwtContract
	sp usecase.SpecializationContract
}

func newRoleRoutes(
	handler *gin.RouterGroup,
	ru usecase.RoleContract,
	j usecase.JwtContract,
	sp usecase.SpecializationContract,
) {
	rr := &roleRoutes{ru: ru, j: j, sp: sp}
	handler.GET("/roles", rr.getAllRoles)
}

type specializationRequestDTO struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type roleRequestDTO struct {
	UUID           string                   `json:"uuid"`
	Name           string                   `json:"name"`
	Specialization specializationRequestDTO `json:"specialization"`
}

type listResponse struct {
	Roles []roleRequestDTO `json:"roles"`
}

func (rr *roleRoutes) getAllRoles(c *gin.Context) {
	access, err := c.Cookie("access")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	_, err = rr.j.CheckToken(access)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	roleList, err := rr.ru.GetAllRoles(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var responseList []roleRequestDTO
	for _, v := range roleList {
		spec, err := rr.sp.GetSpecializationByUUID(c.Request.Context(), v.SpecializationUUID)
		if err != nil {
			errorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		responseList = append(responseList, roleRequestDTO{
			UUID: v.UUID.String(),
			Name: v.Name,
			Specialization: specializationRequestDTO{
				UUID:  spec.UUID.String(),
				Name:  spec.Name,
				Value: spec.Value,
			},
		})
	}
	c.JSON(http.StatusOK, listResponse{responseList})
}
