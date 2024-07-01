package api

import (
    "user-management/controllers"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupProjectRouter(router *gin.RouterGroup, db *gorm.DB) {
    projectController := controllers.NewProjectController(db)

	router.POST("/project", projectController.CreateProject)
	router.PUT("/project/:id", projectController.UpdateProject)
	router.GET("/project", projectController.GetAllProjects)
    
}


