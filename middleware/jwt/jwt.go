package jwt

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/IainMcl/go-behind-the-scenes/internal/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 200
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			code = 400
		} else {
			tokenString := authHeader[len("Bearer "):]
			_, err := util.ParseToken(tokenString)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = 20002
				default:
					code = 20001
				}
			}
		}

		if code != 200 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
