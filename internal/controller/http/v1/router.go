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
	c usecase.CategoryContract,
	cs usecase.CompanyContract,
	mg usecase.MessageContract,
) {
	h := handler.Group("/api/v1")

	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	{
		newProjectRouter(h, p)
		newRegisterRoutes(h, s)
		newLoginRoutes(h, j, u, pr)
		newLogoutRouter(h)
		newChatRoutes(h, c, j, pr, mg)
		newCountryRoute(h, cc)
		newCitizenshipRoutes(h, ct)
		newEduspecialitiesRoutes(h, ed)
		newEmploymentsRoutes(h, ep)
		newSpecializationsRoutes(h, sp)
		newUniversitiesRoutes(h, un)
		newCityRoutes(h, st)
		newCategoryRoutes(h, c)
		newCompanyRoutes(h, cs)
		newProfileRoutes(h, pr)
		newChangeRoutes(h, j, u)
		newMessageRoutes(h, j, mg)
	}
}
