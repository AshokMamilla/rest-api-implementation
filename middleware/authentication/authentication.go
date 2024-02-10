package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api-implementation/config"
	L "rest-api-implementation/middleware/logger"
	mdl "rest-api-implementation/models"
	"time"
)

func GenerateToken(user mdl.User) (string, error) {
	// Create the claims
	claims := mdl.Claims{
		UserID: user.ID.String(),
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	L.RaiLog("D", "Claims Are:"+L.PrintStruct(claims), nil)

	// Convert JWT secret to []byte
	jwtSecret := []byte(config.SecretKey)

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	L.RaiLog("D", "Token:"+L.PrintStruct(tokenString), err)
	return tokenString, nil
}
func RefreshToken(c *gin.Context) {
	db, err := config.OpenDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer config.CloseDB(db)
	// Get refresh token from request
	L.RaiLog("D", "In Refresh Token", nil)
	refreshToken := c.GetHeader("Authorization")
	if refreshToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token is missing"})
		c.Abort()
		return
	}
	// Validate and parse refresh token
	token, err := jwt.Parse(refreshToken[len("Bearer "):], func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		c.Abort()
		return
	}
	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		c.Abort()
		return
	}
	// Check if the token is expired
	exp := time.Unix(int64(claims["exp"].(float64)), 0)
	if exp.Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token has expired"})
		c.Abort()
		return
	}
	// Get user ID from claims
	userID := claims["user_id"].(string)
	// Retrieve user from database based on user ID
	var user mdl.User
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		c.Abort()
		return
	}
	// Generate new access token
	accessToken, err := GenerateToken(user)
	if err != nil {
		L.RaiLog("D", "Error Generating Access Token:\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		c.Abort()
		return
	}
	// Update new access token in the database
	var existingToken mdl.Token
	if err := db.Where("user_id = ?", user.ID).First(&existingToken).Error; err != nil {
		// Handle error (e.g., token not found)
		L.RaiLog("D", "Error finding token in DB:\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find token in DB"})
		c.Abort()
		return
	}
	// Update token value
	existingToken.Token = accessToken

	// Update token expiry
	existingToken.ExpiredAt = time.Now().Add(time.Hour * 1)

	if err := db.Save(&existingToken).Error; err != nil {
		// Handle error (e.g., update failed)
		L.RaiLog("D", "Error updating token in DB:\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update token in DB"})
		c.Abort()
		return
	}
	// Return new access token
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
