package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"lg/internal/usecase"
	"net/http"
)

type projectRouter struct {
	p usecase.ProjectContract
}

type projectListResponse struct {
	Projects []projectDTO `json:"projects"`
}

type projectDTO struct {
	UUID             uuid.UUID `json:"uuid"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	ProjectLink      string    `json:"link"`
	PresentationLink string    `json:"presentation"`
	CreatorID        int64     `json:"creator_id"`
}

func newProjectRouter(handler *gin.RouterGroup, p usecase.ProjectContract) {
	pr := &projectRouter{p: p}
	handler.GET("/project", pr.getAllProjects)
	handler.GET("/project/:uuid", pr.getProjectByName)
	handler.POST("/project", pr.createProject)
	handler.PUT("/project/:uuid", pr.updateProject)
	handler.DELETE("/project/:uuid", pr.deleteProject)
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

func (pr *projectRouter) getProjectByName(c *gin.Context) {
	projectKey, err := uuid.FromString(c.Param("uuid"))
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
	name, err := pr.p.CreateProject(c.Request.Context(), projectToEntity(*req))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("location", fmt.Sprintf("/api/v1/project/%s", name))
	c.JSON(http.StatusCreated, nil)
}

func (pr *projectRouter) updateProject(c *gin.Context) {
	projectKey, err := uuid.FromString(c.Param("uuid"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	req := new(projectDTO)
	if err := c.ShouldBindJSON(req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	projectEntity := projectToEntity(*req)
	projectEntity.UUID = projectKey
	err = pr.p.UpdateProjectByUUID(c.Request.Context(), projectEntity)
	if err != nil {
		errorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (pr *projectRouter) deleteProject(c *gin.Context) {
	projectKey, err := uuid.FromString(c.Param("uuid"))
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
