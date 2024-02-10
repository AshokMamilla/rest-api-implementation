package usersignup

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"rest-api-implementation/config"
	auth "rest-api-implementation/middleware/authentication"
	pwd "rest-api-implementation/middleware/hashpassword"
	L "rest-api-implementation/middleware/logger"
	vld "rest-api-implementation/middleware/validations"
	mdl "rest-api-implementation/models"
	"time"
)

var db *gorm.DB

func SignUpService(c *gin.Context) {
	var newUser mdl.User
	defer config.CloseDB(db)
	L.RaiLog("D", "Opening database Connection", nil)
	db, err := config.OpenDB()
	//validate request
	if err := c.ShouldBindJSON(&newUser); err != nil {
		L.RaiLog("E", "Error Processing the Request", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validate email format
	if !vld.IsValidEmail(newUser.Email) {
		L.RaiLog("E", "Here User Sent Invalid Email format :"+L.PrintStruct(newUser.Email), nil)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}
	// Check if email already exists
	var existingUser mdl.User
	result := db.Where("email = ?", newUser.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		L.RaiLog("E", "Email is already registered"+L.PrintStruct(newUser.Email), nil)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}
	if result.RowsAffected > 0 {
		// Record already exists, return error
		L.RaiLog("E", "Email is already registered: "+existingUser.Email, nil)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}
	// Hash the password before saving it
	hashedPassword, err := pwd.HashPassword(newUser.Password)
	if err != nil {
		L.RaiLog("E", "Error generating hash password:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	// Create a new UUID for the user
	newUser.ID = uuid.New()
	newUser.Password = hashedPassword
	// Create the user in the database
	if err := db.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
func SignInService(c *gin.Context) {
	var signInReq mdl.SignInRequest
	// Bind JSON request body into signInReq struct
	if err := c.ShouldBindJSON(&signInReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	defer config.CloseDB(db)
	L.RaiLog("D", "Opening database Connection", nil)
	db, err := config.OpenDB()

	// Retrieve user from the database
	var user mdl.User
	if err := db.Where("email = ?", signInReq.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare password
	err = pwd.ComparePasswords(user.Password, signInReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user)
	if err != nil {
		L.RaiLog("D", "Error Generating Token:\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Store the token with issued time and expiry time
	newToken := mdl.Token{
		UserID:    user.ID,
		Email:     user.Email,
		Token:     token,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Hour * 1),
	}

	if err := db.Create(&newToken).Error; err != nil {
		L.RaiLog("D", "Error Storing Token:\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store token"})
		return
	}
	// Return JWT token as response
	c.JSON(http.StatusOK, gin.H{"token": token})
}
