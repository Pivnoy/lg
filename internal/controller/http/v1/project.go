package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lg/internal/entity"
	"lg/internal/usecase"
	"net/http"
)

type projectRouter struct {
	p usecase.ProjectContract
}

func newProjectRouter(handler *gin.RouterGroup, p usecase.ProjectContract) {
	pr := &projectRouter{p: p}
	handler.GET("/project", pr.getAllProjects)
	handler.GET("/project/:name", pr.getProjectByName)
}

func (pr *projectRouter) getAllProjects(c *gin.Context) {
	projectList, err := pr.p.GetAllProjects(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var response []entity.Project
	if projectList != nil {
		response = projectList
	} else {
		response = make([]entity.Project, 0)
	}
	c.JSON(http.StatusOK, response)
}

func (pr *projectRouter) getProjectByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		errorResponse(c, http.StatusBadRequest, "Project name cannot be empty")
		return
	}
	project, err := pr.p.GetProjectByName(c.Request.Context(), name)
	if err != nil {
		errorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, project)
}

type projectRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Link         string `json:"link"`
	Presentation string `json:"presentation"`
}

func (pr *projectRouter) createProject(c *gin.Context) {
	req := new(projectRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	name, err := pr.p.CreateProject(c.Request.Context(), entity.Project{
		Name:         req.Name,
		Description:  req.Description,
		Link:         req.Link,
		Presentation: req.Presentation,
	})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("location", fmt.Sprintf("/api/v1/project/%s", name))
	c.JSON(http.StatusCreated, nil)
}
