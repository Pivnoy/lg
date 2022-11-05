package v1

import (
	"github.com/gin-gonic/gin"
	"lg/internal/usecase"
	"net/http"
)

type profileRoutes struct {
	ps usecase.ProfileContract
	j  usecase.JwtContract
}

type profileRequestDTO struct {
	UserUUID           string `json:"userUuid"`
	Firstname          string `json:"firstname"`
	Lastname           string `json:"lastname"`
	Patronymic         string `json:"patronymic"`
	CountryUUID        string `json:"countryUuid"`
	CityUUID           string `json:"cityUuid"`
	CitizenshipUUID    string `json:"citizenshipUuid"`
	Gender             string `json:"gender"`
	Phone              string `json:"phone"`
	Email              string `json:"email"`
	UniversityUUID     string `json:"universityUuid"`
	EduspecialityUUID  string `json:"eduspecialityUuid"`
	GraduationYear     uint   `json:"graduationYear"`
	EmploymentUUID     string `json:"employmentUuid"`
	Experience         uint   `json:"experience"`
	AchievementUUID    string `json:"achievementUuid"`
	TeamUUID           string `json:"teamUuid"`
	SpecializationUUID string `json:"specializationUuid"`
	CompanyInn         string `json:"companyInn"`
	CompanyName        string `json:"companyName"`
}

type profileResponseDTO struct {
	UUID       string `json:"uuid"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Patronymic string `json:"patronymic"`
}

func newProfileRoutes(handler *gin.RouterGroup, ps usecase.ProfileContract, j usecase.JwtContract) {
	pr := &profileRoutes{ps: ps, j: j}
	handler.POST("/profile", pr.createProfile)
}

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
	newProject, err := pr.ps.CreateProfile(c.Request.Context(), profileEntity, req.CompanyName, req.CompanyInn)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response := profileToDTO(newProject)
	c.JSON(http.StatusCreated, response)
}
