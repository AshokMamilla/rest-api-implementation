package products

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	Au "rest-api-implementation/middleware/authorization"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/protected", Au.AuthorizeRequest(), ProductHandler)
}

// ProductHandler  Protected route handler
func ProductHandler(c *gin.Context) {
	// Access user information from context
	user := c.MustGet("user").(jwt.MapClaims)
	c.JSON(http.StatusOK, gin.H{"message": "Welcome," + user["email"].(string)})
}
