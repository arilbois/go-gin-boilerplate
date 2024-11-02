package controllers

import (
	"msvc-syahril-app/app/helpers"
	"msvc-syahril-app/app/models"
	"msvc-syahril-app/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
    service *services.ProjectService
}

func NewProjectController(service *services.ProjectService) *ProjectController {
    return &ProjectController{
        service: service,
    }
}

// GetAllProjects godoc
// @Summary Get all projects
// @Description Retrieve all available projects
// @Tags projects
// @Produce json
// @Success 200 {object} helpers.APIResponse
// @Failure 500 {object} helpers.APIResponse
// @Router /projects [get]
func (pc *ProjectController) GetAllProjects(c *gin.Context) {
    projects, err := pc.service.GetAllProjects()
    if err != nil {
        response := helpers.ErrorResponse(http.StatusInternalServerError, "Failed to retrieve projects", nil)
        c.JSON(http.StatusInternalServerError, response)
        return
    }
    response := helpers.SuccessResponse("Projects retrieved successfully", projects)
    c.JSON(http.StatusOK, response)
}

// GetProjectByID godoc
// @Summary Get project by ID
// @Description Retrieve project details by ID
// @Tags projects
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} helpers.APIResponse
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /projects/{id} [get]
func (pc *ProjectController) GetProjectByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        response := helpers.ErrorResponse(http.StatusBadRequest, "Invalid project ID", nil)
        c.JSON(http.StatusBadRequest, response)
        return
    }

    project, err := pc.service.GetProjectByID(id)
    if err != nil {
        response := helpers.ErrorResponse(http.StatusNotFound, "Project not found", nil)
        c.JSON(http.StatusNotFound, response)
        return
    }

    response := helpers.SuccessResponse("Project retrieved successfully", project)
    c.JSON(http.StatusOK, response)
}

// CreateProject godoc
// @Summary Create a new project
// @Description Create a new project with the provided data
// @Tags projects
// @Accept json
// @Produce json
// @Param project body models.Project true "Project"
// @Success 201 {object} helpers.APIResponse
// @Failure 400 {object} helpers.APIResponse
// @Failure 500 {object} helpers.APIResponse
// @Router /projects [post]
func (pc *ProjectController) CreateProject(c *gin.Context) {
    var project models.Project
    if err := c.ShouldBindJSON(&project); err != nil {
        response := helpers.ErrorResponse(http.StatusBadRequest, err.Error(), nil)
        c.JSON(http.StatusBadRequest, response)
        return
    }

    if err := pc.service.CreateProject(&project); err != nil {
        response := helpers.ErrorResponse(http.StatusInternalServerError, "Failed to create project", nil)
        c.JSON(http.StatusInternalServerError, response)
        return
    }

    response := helpers.SuccessResponse("Project created successfully", project)
    c.JSON(http.StatusCreated, response)
}

// UpdateProject godoc
// @Summary Update a project
// @Description Update project data by ID
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Param project body models.Project true "Project"
// @Success 200 {object} helpers.APIResponse
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Failure 500 {object} helpers.APIResponse
// @Router /projects/{id} [put]
func (pc *ProjectController) UpdateProject(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        response := helpers.ErrorResponse(http.StatusBadRequest, "Invalid project ID", nil)
        c.JSON(http.StatusBadRequest, response)
        return
    }

    var project models.Project
    if err := c.ShouldBindJSON(&project); err != nil {
        response := helpers.ErrorResponse(http.StatusBadRequest, err.Error(), nil)
        c.JSON(http.StatusBadRequest, response)
        return
    }

    project.ID = id

    if err := pc.service.UpdateProject(&project); err != nil {
        response := helpers.ErrorResponse(http.StatusInternalServerError, "Failed to update project", nil)
        c.JSON(http.StatusInternalServerError, response)
        return
    }

    response := helpers.SuccessResponse("Project updated successfully", project)
    c.JSON(http.StatusOK, response)
}

// DeleteProject godoc
// @Summary Delete a project
// @Description Delete a project by ID
// @Tags projects
// @Param id path int true "Project ID"
// @Success 200 {object} helpers.APIResponse
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /projects/{id} [delete]
func (pc *ProjectController) DeleteProject(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        response := helpers.ErrorResponse(http.StatusBadRequest, "Invalid project ID", nil)
        c.JSON(http.StatusBadRequest, response)
        return
    }

    if _, err := pc.service.GetProjectByID(id); err != nil {
        response := helpers.ErrorResponse(http.StatusNotFound, "Project not found", nil)
        c.JSON(http.StatusNotFound, response)
        return
    }

    if err := pc.service.DeleteProject(id); err != nil {
        response := helpers.ErrorResponse(http.StatusInternalServerError, "Failed to delete project", nil)
        c.JSON(http.StatusInternalServerError, response)
        return
    }

    response := helpers.SuccessResponse("Project deleted successfully", nil)
    c.JSON(http.StatusOK, response)
}
