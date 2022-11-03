package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"lg/internal/usecase"
	"net/http"
)

type projectRouter struct {
	p usecase.ProjectContract
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

func newProjectRouter(handler *gin.RouterGroup, p usecase.ProjectContract) {
	pr := &projectRouter{p: p}
	handler.GET("/project", pr.getAllProjects)
	handler.GET("/project/:uuid", pr.getProjectByUUID)
	handler.POST("/project", pr.createProject)
	handler.DELETE("/project/:uuid", pr.deleteProjectByUUID)
}

func (pr *projectRouter) getAllProjects(c *gin.Context) {
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

func (pr *projectRouter) getProjectByUUID(c *gin.Context) {
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

func (pr *projectRouter) createProject(c *gin.Context) {
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

func (pr *projectRouter) deleteProjectByUUID(c *gin.Context) {
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
