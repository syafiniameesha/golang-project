package controllers

import (
    "net/http"
    "strconv"
    "errors"
    "user-management/models"
    "user-management/helpers"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type UserController struct {
    BaseController
}

func NewUserController(db *gorm.DB) *UserController {
    return &UserController{
        BaseController: BaseController{DB: db},
    }
}

type CreateUserInput struct {
    FirstName string `json:"firstname" binding:"required"`
    LastName  string `json:"lastname" binding:"required"`
    Email     string `json:"email" binding:"required,email"`
    Password  string `json:"password" binding:"required,min=6"`
}

func (uc *UserController) CreateUser(c *gin.Context) {
    var input CreateUserInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := models.User{
        FirstName: input.FirstName,
        LastName: input.LastName,
        Email:    input.Email,
        Password: input.Password,
    }

    if err := uc.Create(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
    // Get the user ID from the URL parameter
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    // Bind the new data from the request body
    var input CreateUserInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Find the user in the database
    var user models.User
    if err := uc.FindByID(&user, id); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Update the user fields
    user.FirstName = input.FirstName
    user.LastName = input.LastName
    user.Email = input.Email
    user.Password = input.Password

    // Save the updated user to the database
    if err := uc.Update(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
}

// GetUserDetails fetches user details based on token
func (uc *UserController) GetUserDetails(c *gin.Context) {
    userID, err := getUserIdFromToken(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    // Fetch user from DB
    var user models.User
    if err := uc.DB.Where("id = ?", userID).First(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": user})
}

func getUserIdFromToken(c *gin.Context) (uint, error) {
    tokenString := c.GetHeader("Authorization")
    if tokenString == "" {
        return 0, errors.New("Authorization header is required")
    }

    token, err := helpers.VerifyToken(tokenString)
    if err != nil {
        return 0, err
    }

    return token.UserID, nil
}