package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"lg/internal/usecase"
	"net/http"
)

type profileRoutes struct {
	ps usecase.ProfileContract
	j  usecase.JwtContract
}

type profileRequestDTO struct {
	UserUUID           string `json:"userUuid"`           // string -> uuid.UUID
	Firstname          string `json:"firstname"`          // simple string
	Lastname           string `json:"lastname"`           // simple string
	Patronymic         string `json:"patronymic"`         // simple string
	CountryUUID        string `json:"countryUuid"`        // string -> uuid.UUID
	CityUUID           string `json:"cityUuid"`           // string -> uuid.UUID
	CitizenshipUUID    string `json:"citizenshipUuid"`    // string -> uuid.UUID
	Gender             string `json:"gender"`             // simple string
	Phone              string `json:"phone"`              // simple string
	Email              string `json:"email"`              // simple string
	UniversityUUID     string `json:"universityUuid"`     // string ""-> uuid.UUID (null)
	EduspecialityUUID  string `json:"eduspecialityUuid"`  // string ""-> uuid.UUID (null)
	GraduationYear     string `json:"graduationYear"`     // string -> int
	EmploymentUUID     string `json:"employmentUuid"`     // string -> uuid.UUID
	Experience         string `json:"experience"`         // string -> int
	AchievementUUID    string `json:"achievement"`        // text -> create achivment and add uuid
	SpecializationUUID string `json:"specializationUuid"` // string -> uuid.UUID
	CompanyInn         string `json:"companyInn"`         //
	CompanyName        string `json:"companyName"`        // все также
	// skills: - массив строк
}

//type profileResponseDTO struct {
//}

type profileResponseDTO struct {
	UUID       string `json:"uuid"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Patronymic string `json:"patronymic"`
}

func newProfileRoutes(handler *gin.RouterGroup, ps usecase.ProfileContract, j usecase.JwtContract) {
	pr := &profileRoutes{ps: ps, j: j}
	handler.POST("/profile", pr.createProfile)
	handler.GET("/profile/:uuid", pr.getProfileByUser)
}

// FIXME swagger
// @Summary CreateProfile
// @Tags Profile
// @Description Create profile
// @Param input body profileRequestDTO true "enter info project"
// @Success 201 {object} profileResponseDTO
// @Failure 400 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/profile [post]
func (pr *profileRoutes) createProfile(c *gin.Context) {
	access, err := c.Cookie("access")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	_, err = pr.j.CheckToken(access)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	req := new(profileRequestDTO)
	if err := c.ShouldBindJSON(req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	profileEntity, err := profileToEntity(*req)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newProject, err := pr.ps.CreateProfile(c.Request.Context(), profileEntity, req.CompanyName, req.CompanyInn, req.AchievementUUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response := profileToDTO(newProject)
	c.JSON(http.StatusCreated, response)
}

// TODO доделать профиль, как пахан фиксанет базу
func (pr *profileRoutes) getProfileByUser(c *gin.Context) {
	access, err := c.Cookie("access")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	_, err = pr.j.CheckToken(access)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userKey, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err = pr.ps.GetProfileByUser(c.Request.Context(), userKey)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

}
