package main

import (
	"github.com/gin-gonic/gin"
	"rest-api-implementation/api/products"
	"rest-api-implementation/api/users"
	C "rest-api-implementation/config"
	L "rest-api-implementation/middleware/logger"
)

func main() {
	// Initialize Gin
	r := gin.Default()
	/*
	  InitDB function performs only Models Auto Migration it doesn't
	  provide live db connection.
	  To Avoid DB Connection Pooling issues, Each controller explicitly
	  open db connections and closes once they serve each Api Request.
	*/
	_, err := C.InitDB()
	if err != nil {
		L.RaiLog("E", "Error Occured while connecting database", err)
	}

	//User Sign-up Route
	users.SignUpRoute(r)
	//User Sign-in Route
	users.SignINRoute(r)
	//users RefreshToken route
	users.RefreshTokenRoute(r)
	//users RevokeToken route
	users.RovokeTokenRoute(r)
	//Product Route
	products.ProductRoutes(r)
	// Run the server
	err = r.Run(":8080")
	if err != nil {
		L.RaiLog("E", "Error while dealing with server port", err)
		return
	}
}
