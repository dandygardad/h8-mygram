package middlewares

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"mygram/helper"
	"net/http"
	"os"
	"strings"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		inputAuth := c.Request.Header.Get("Authorization")
		hasBearer := strings.HasPrefix(inputAuth, "Bearer")

		if !hasBearer {
			helper.ResponseError(c, "unauthorized", http.StatusUnauthorized)
			return
		}

		// Split token
		strToken := strings.Split(inputAuth, " ")
		if len(strToken) == 1 {
			helper.ResponseError(c, "unauthorized", http.StatusUnauthorized)
			return
		}

		// Check token
		token, err := jwt.Parse(strToken[1], func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unauthorized")
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			helper.ResponseError(c, "unauthorized", http.StatusUnauthorized)
			return
		}

		// validate
		result, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			helper.ResponseError(c, "unauthorized", http.StatusUnauthorized)
			return
		}

		c.Set("userData", result)
		c.Next()
	}
}
