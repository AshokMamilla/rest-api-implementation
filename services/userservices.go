package usersignup

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"rest-api-implementation/config"
	pwd "rest-api-implementation/middleware/hashpassword"
	L "rest-api-implementation/middleware/logger"
	vld "rest-api-implementation/middleware/validations"
	mdl "rest-api-implementation/models"
)

var db *gorm.DB

func SignUpService(c *gin.Context) {
	var newUser mdl.User
	defer config.CloseDB(db)
	L.RaiLog("D", "Opening database Connection", nil)
	db, err := config.OpenDB()
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
