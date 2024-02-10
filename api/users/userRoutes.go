package users

import (
	"github.com/gin-gonic/gin"
	auth "rest-api-implementation/middleware/authentication"
	srv "rest-api-implementation/services"
)

func SignUpRoute(r *gin.Engine) {
	// Define signup handler with closure
	r.POST("/signup", func(c *gin.Context) {
		srv.SignUpService(c)
	})

}
func SignINRoute(r *gin.Engine) {
	// Define signup handler with closure
	r.POST("/signin", func(c *gin.Context) {
		srv.SignInService(c)
	})

}
func RefreshTokenRoute(r *gin.Engine) {
	r.POST("/refresh-token", func(c *gin.Context) { auth.RefreshToken(c) })
}
func RovokeTokenRoute(r *gin.Engine) {
	r.POST("/revoke-token", func(c *gin.Context) { auth.RevokeToken(c) })
}
