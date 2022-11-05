package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"lg/internal/usecase"
)

func NewRouter(handler *gin.Engine,
	p usecase.ProjectContract,
	s usecase.RegisterContract,
	j usecase.JwtContract,
	u usecase.UserContract,
	pr usecase.ProfileContract,
	c usecase.ChatContract,
	cc usecase.CountryContract,
	ct usecase.CitizenshipContract,
	ed usecase.EduspecialityContract,
	ep usecase.EmploymentContract,
	sp usecase.SpecializationContract,
	un usecase.UniversityContract,
	st usecase.CityContract,
	ccc usecase.CategoryContract,
	cs usecase.CompanyContract,
	mg usecase.MessageContract,
) {
	h := handler.Group("/api/v1")

	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	{
		newProjectRouter(h, p, j)
		newRegisterRoutes(h, s)
		newLoginRoutes(h, j, u, pr)
		newLogoutRouter(h)
		newChatRoutes(h, c, j, pr, mg)
		newCountryRoute(h, cc, j)
		newCitizenshipRoutes(h, ct, j)
		newEduspecialitiesRoutes(h, ed, j)
		newEmploymentsRoutes(h, ep, j)
		newSpecializationsRoutes(h, sp, j)
		newUniversitiesRoutes(h, un, j)
		newCityRoutes(h, st, j)
		newCategoryRoutes(h, ccc, j)
		newCompanyRoutes(h, cs, j)
		newProfileRoutes(h, pr, j)
		newChangeRoutes(h, j, u)
		newMessageRoutes(h, j, mg, c)
	}
}
