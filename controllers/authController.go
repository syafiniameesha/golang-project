package controllers

import (
    "time"
    "log"
    "net/http"
    "user-management/models"
    "user-management/helpers"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

type AuthController struct {
    DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
    return &AuthController{DB: db}
}

type SignupInput struct {
    FirstName string `json:"firstname" binding:"required"`
    LastName  string `json:"lastname" binding:"required"`
    Email     string `json:"email" binding:"required,email"`
    Password  string `json:"password" binding:"required,min=6"`
}

type SigninInput struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

type ForgotPasswordInput struct {
    Email string `json:"email" binding:"required,email"`
}

type ResetPasswordInput struct {
    // Token    string `json:"token" binding:"required"`
    Password string `json:"password" binding:"required,min=6"`
}

func (ac *AuthController) Signup(c *gin.Context) {
    var input SignupInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Check if user already exists
    var existingUser models.User
    if err := ac.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
        return
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    // Create user record
    newUser := models.User{
        FirstName: input.FirstName,
        LastName:  input.LastName,
        Email:     input.Email,
        Password:  string(hashedPassword),
    }

    if err := ac.DB.Create(&newUser).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    // Generate JWT tokens
    accessToken, err := helpers.GenerateAccessToken(newUser.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
        return
    }

    refreshToken, err := helpers.GenerateRefreshToken(newUser.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
        return
    }

    // Update user with tokens
    newUser.Token = accessToken
    newUser.RefreshToken = refreshToken
    if err := ac.DB.Save(&newUser).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save tokens"})
        return
    }

    // Return response with user and tokens
    c.JSON(http.StatusOK, gin.H{
        "user":           newUser,
        "access_token":   accessToken,
        "refresh_token":  refreshToken,
    })
}

func (ac *AuthController) Login(c *gin.Context) {
    var input SigninInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    // Find user by email
    var user models.User
    if err := ac.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Compare hashed password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Generate JWT access token
    accessToken, err := helpers.GenerateAccessToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
        return
    }

    // Update user's access token in database (optional, depends on your application logic)
    user.Token = accessToken
    if err := ac.DB.Save(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": accessToken})
}

func (ac *AuthController) ForgotPassword(c *gin.Context) {
    var input ForgotPasswordInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Find user by email
    var user models.User
    if err := ac.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Generate password reset token
    resetToken, err := helpers.GeneratePasswordResetToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate password reset token"})
        return
    }

    // Set password reset token and expiry in user record
    user.PasswordResetToken = resetToken
    expiresAt := time.Now().Add(time.Hour)
    user.PasswordResetExpiresAt = &expiresAt 

    if err := ac.DB.Save(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save password reset token"})
        return
    }

    // Send password reset email to user
    resetLink := "http://example.com/reset-password?token=" + resetToken
    emailBody := "Here is the link to reset your password: <a href=\"" + resetLink + "\">Reset Password</a>"
    if err := helpers.SendEmail(user.Email, "Password Reset", emailBody); err != nil {
        log.Printf("Failed to send email to %s: %v", user.Email, err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Password reset instructions sent to your email"})
}

func (ac *AuthController) ResetPassword(c *gin.Context) {
    // Retrieve token from URL parameter
    token := c.Param("token")

    // Find user by password reset token
    var user models.User
    if err := ac.DB.Where("password_reset_token = ? AND password_reset_expires_at > ?", token, time.Now()).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired password reset token"})
        return
    }

    // Bind JSON input (if any other data is needed)
    var input ResetPasswordInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash the new password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    // Update user's password and clear reset token
    user.Password = string(hashedPassword)
    user.PasswordResetToken = ""
    user.PasswordResetExpiresAt = nil // Set to NULL

    if err := ac.DB.Save(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset password"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Password reset successful"})
}

func (ac *AuthController) Logout(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
