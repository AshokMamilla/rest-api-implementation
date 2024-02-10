package users

import (
	"github.com/gin-gonic/gin"
	L "rest-api-implementation/middleware/logger"
	srv "rest-api-implementation/services"
)

func SignUpRoute(r *gin.Engine) {

	// Define signup handler with closure
	r.POST("/signup", func(c *gin.Context) {
		srv.SignUpService(c)
	})
	// Run the server
	err := r.Run(":8080")
	if err != nil {
		L.RaiLog("E", "Error while dealing with server port", err)
		return
	}
}
