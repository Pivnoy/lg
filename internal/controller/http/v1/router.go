package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
)

func NewRouter(handler *gin.Engine,
	p usecase.ProjectContract,
	s usecase.RegisterContract,
	j usecase.JwtContract,
	u usecase.UserContract,
	pr usecase.ProfileContract,
	cc usecase.CountryContract,
	ct usecase.CitizenshipContract,
	ed usecase.EduspecialityContract,
	ep usecase.EmploymentContract,
	sp usecase.SpecializationContract,
	un usecase.UniversityContract,
	st usecase.CityContract,
) {
	h := handler.Group("/api/v1")
	{
		newProjectRouter(h, p)
		newRegisterRoutes(h, s)
		newLoginRoutes(h, j, u, pr)
		newLogoutRouter(h)
		newCountryRoute(h, cc)
		newCitizenshipRoutes(h, ct)
		newEduspecialitiesRoutes(h, ed)
		newEmploymentsRoutes(h, ep)
		newSpecializationsRoutes(h, sp)
		newUniversitiesRoutes(h, un)
		NewCityRoutes(h, st)
	}
}
