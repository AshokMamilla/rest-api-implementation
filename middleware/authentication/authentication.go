package authentication

import (
	"github.com/dgrijalva/jwt-go"
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
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	//L.RaiLog("D", "Claims Are:"+L.PrintStruct(claims), nil)
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
