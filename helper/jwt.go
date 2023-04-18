package helper

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"mygram/model/entity"
	"os"
)

type UserJWT struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func GenerateToken(user entity.User) string {
	claims := jwt.MapClaims{
		"id":       user.Id,
		"email":    user.Email,
		"username": user.Username,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := parseToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	return signedString
}
