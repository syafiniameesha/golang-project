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
    Password string `json:"password" binding:"required,min=6"`
    ConfirmPassword string `json:"confirmPassword" binding:"required,min=6"`
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

    var user models.User
    if err := ac.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, err := helpers.GenerateAccessToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token, "user": user})
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
    expiresAt := time.Now().Add(9 * time.Hour)
    log.Printf("Current time 3: %s", expiresAt.Format("2006-01-02 15:04:05"))

    user.PasswordResetExpiresAt = &expiresAt 

    if err := ac.DB.Save(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save password reset token"})
        return
    }

    // Send password reset email to user
    resetLink := "http://localhost:5173/resetPassword?token=" + resetToken
    emailBody := "Here is the link to reset your password: <a href=\"" + resetLink + "\">Reset Password</a>"
    if err := helpers.SendEmail(user.Email, "Password Reset", emailBody); err != nil {
        log.Printf("Failed to send email to %s: %v", user.Email, err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Password reset instructions sent to your email", "email": user.Email})
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

    // Bind JSON input
    var input ResetPasswordInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Check if passwords match
    if input.Password != input.ConfirmPassword {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
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
