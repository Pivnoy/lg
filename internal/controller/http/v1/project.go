package v1

import (
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
