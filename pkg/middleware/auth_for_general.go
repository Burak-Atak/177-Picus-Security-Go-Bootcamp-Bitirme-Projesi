package middleware

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/helpers"
	jwtHelper "github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthForGeneral is a middleware function that checks for valid jwt token in the request header.
func AuthForGeneral() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.GetHeader("Authorization") != "" {
			_, err := jwtHelper.VerifyToken(context.GetHeader("Authorization"))
			if err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
				context.Abort()
				return

			} else {
				context.Next()
				context.Abort()
				return
			}
		} else {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": helpers.UnAuthorizedError.Error(),
			})
			context.Abort()
			return
		}
	}
}
