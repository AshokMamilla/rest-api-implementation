package authentication

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api-implementation/config"
	mdl "rest-api-implementation/models"
	"time"
)

// RevokeToken revokes the provided token and adds it to the blacklist
func RevokeToken(c *gin.Context) {
	// Get token to revoke from request
	token := c.GetHeader("Authorization")
	if token == "" || token == "Bearer <access_token>" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token to revoke is missing"})
		return
	}
	// Add token to blacklist
	blacklistEntry := mdl.TokenBlacklist{
		Token:     token,
		Reason:    "Revoked by user",
		ExpiresAt: time.Now().AddDate(0, 0, 7),
		CreatedAt: time.Now(),
	}

	// Get database connection
	db, err := config.OpenDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer config.CloseDB(db)
	// Save token to blacklist in the database
	if err := db.Create(&blacklistEntry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add token to blacklist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Token revoked successfully"})
}
