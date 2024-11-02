package routes

import (
	"msvc-syahril-app/app/controllers"
	"msvc-syahril-app/app/repositories"
	"msvc-syahril-app/app/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    projectRepo := repositories.NewProjectRepository()
    projectService := services.NewProjectService(projectRepo)
    projectController := controllers.NewProjectController(projectService)

    router.GET("/projects", projectController.GetAllProjects)
    router.GET("/projects/:id", projectController.GetProjectByID)
    router.POST("/projects", projectController.CreateProject)
    router.PUT("/projects/:id", projectController.UpdateProject)
    router.DELETE("/projects/:id", projectController.DeleteProject)

    return router
}
