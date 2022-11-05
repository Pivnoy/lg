package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"lg/internal/entity"
	"lg/internal/usecase"
	"net/http"
	"time"
)

type projectRoutes struct {
	p usecase.ProjectContract
	j usecase.JwtContract
	m usecase.MessageContract
}

type projectListResponse struct {
	Projects []projectDTO `json:"projects"`
}

type responseUUID struct {
	UUID uuid.UUID `json:"uuid"`
}

type projectDTO struct {
	UUID             uuid.UUID `json:"uuid"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	CategoryUUID     uuid.UUID `json:"category_uuid"`
	ProjectLink      string    `json:"project_link"`
	PresentationLink string    `json:"presentation_link"`
	CreatorUUID      uuid.UUID `json:"creator_uuid"`
	IsVisible        string    `json:"is_visible"`
}

func newProjectRouter(handler *gin.RouterGroup, p usecase.ProjectContract, j usecase.JwtContract, m usecase.MessageContract) {
	pr := &projectRoutes{p: p, j: j, m: m}

	handler.GET("/project", pr.getAllProjects)
	handler.GET("/project/:uuid", pr.getProjectByUUID)
	handler.POST("/project", pr.createProject)
	handler.DELETE("/project/:uuid", pr.deleteProjectByUUID)
	handler.POST("/ebaloPsa", pr.acceptOrRejectToProject)
}

// @Summary GetAllProjects
// @Tags Projects
// @Description Get all projects
// @Success 200 {object} projectListResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/project [get]
func (pr *projectRoutes) getAllProjects(c *gin.Context) {
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
	projectList, err := pr.p.GetAllProjects(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var responseList []projectDTO
	for _, v := range projectList {
		responseList = append(responseList, projectToDTO(v))
	}
	c.JSON(http.StatusOK, projectListResponse{responseList})
}

// @Summary GetProjectByUUID
// @Tags Projects
// @Description Get project by UUID
// @Param uuid path string true "Enter id book"
// @Success 200 {object} projectDTO
// @Failure 400 {object} errResponse
// @Failure 404 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/project/{uuid} [get]
func (pr *projectRoutes) getProjectByUUID(c *gin.Context) {
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
	projectKey, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	project, err := pr.p.GetProjectByUUID(c.Request.Context(), projectKey)
	if err != nil {
		errorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, projectToDTO(project))
}

// @Summary CreateProject
// @Tags Projects
// @Description Create project
// @Param input body projectDTO true "enter info project"
// @Success 201 {object} responseUUID
// @Failure 400 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/project [post]
func (pr *projectRoutes) createProject(c *gin.Context) {
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
	req := new(projectDTO)
	if err := c.ShouldBindJSON(req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	projectKey, err := pr.p.CreateProject(c.Request.Context(), projectToEntity(*req))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, responseUUID{UUID: projectKey})
}

func (pr *projectRoutes) deleteProjectByUUID(c *gin.Context) {
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
	projectKey, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = pr.p.DeleteProjectByUUID(c.Request.Context(), projectKey)
	if err != nil {
		errorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

type acceptOrRejectRequest struct {
	AuthorUUID string `json:"authorUUID"`
	ChatUUID   string `json:"chatUUID"`
	Decision   string `json:"decision"`
}

func (pr *projectRoutes) acceptOrRejectToProject(c *gin.Context) {
	access, err := c.Cookie("access")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	tokenUUID, err := pr.j.CheckToken(access)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	creatorUUID, err := uuid.Parse(tokenUUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var accept acceptOrRejectRequest
	if err := c.ShouldBindJSON(&accept); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	authorUUID, err := uuid.Parse(accept.AuthorUUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	chatUUID, err := uuid.Parse(accept.ChatUUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = pr.m.UpdateMessageStatus(c.Request.Context(), authorUUID, chatUUID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	msg := entity.Message{
		AuthorUUID:   creatorUUID,
		Type:         "text",
		CreationDate: time.Now(),
		ChatUUID:     chatUUID,
	}
	if accept.Decision == "accept" {
		msg.Content = "Жду в комнате, сосочка :))))"
	} else if accept.Decision == "reject" {
		msg.Content = "Бро, сори, санечка не снимает :((("
	} else {
		errorResponse(c, http.StatusInternalServerError, "cannot get decision from creator")
	}
	err = pr.m.StoreMessage(c.Request.Context(), msg)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
