package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"small-run/fx_gin/shared/tools"
)

func NewJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var msg string
		token := c.Query("token")
		_, err := tools.ParseToken(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				msg = "token超时"
			default:
				msg = "token鉴权失败"
			}

			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": msg,
			})

			c.Abort()

			return
		}

		c.Next()
	}
}
