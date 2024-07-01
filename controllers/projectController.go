package controllers

import (
    "net/http"
    "strconv"
    "time"
    "user-management/models"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type ProjectController struct {
    BaseController
}

func NewProjectController(db *gorm.DB) *ProjectController {
    return &ProjectController{
        BaseController: BaseController{DB: db},
    }
}

type CreateProjectInput struct {
    Prefix     string `json:"prefix" binding:"required"`
    Suffix     string `json:"suffix" binding:"required"`
    FullCode   string `json:"fullCode" binding:"required"`
    Name       string `json:"name" binding:"required"`
    StartDate  string `json:"startDate" binding:"required"`
    EndDate    string `json:"endDate" binding:"required"`
    StatusID   uint   `json:"statusId" binding:"required"`
    TypeID     uint   `json:"typeId" binding:"required"`
    UserID     uint   `json:"userId" binding:"required"`
}

func (pc *ProjectController) CreateProject(c *gin.Context) {
    var input CreateProjectInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Parse the start and end dates from strings to time.Time
    startDate, err := time.Parse("2006-01-02", input.StartDate)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
        return
    }
    endDate, err := time.Parse("2006-01-02", input.EndDate)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format"})
        return
    }

    project := models.Project{
        Prefix:    input.Prefix,
        Suffix:    input.Suffix,
        FullCode:  input.FullCode,
        Name:      input.Name,
        StartDate: startDate,
        EndDate:   endDate,
        StatusID:  input.StatusID,
        TypeID:    input.TypeID,
        UserID:    input.UserID,
    }

    if err := pc.DB.Create(&project).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": project})
}

func (pc *ProjectController) UpdateProject(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
        return
    }

    var input CreateProjectInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var project models.Project
    if err := pc.DB.First(&project, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
        return
    }

    // Parse the start and end dates from strings to time.Time
    startDate, err := time.Parse("2006-01-02", input.StartDate)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
        return
    }
    endDate, err := time.Parse("2006-01-02", input.EndDate)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format"})
        return
    }

    project.Prefix = input.Prefix
    project.Suffix = input.Suffix
    project.FullCode = input.FullCode
    project.Name = input.Name
    project.StartDate = startDate
    project.EndDate = endDate
    project.StatusID = input.StatusID
    project.TypeID = input.TypeID
    project.UserID = input.UserID

    if err := pc.DB.Save(&project).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": project})
}

func (pc *ProjectController) GetAllProjects(c *gin.Context) {
    var projects []models.Project
    if err := pc.DB.Find(&projects).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve projects"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": projects})
}