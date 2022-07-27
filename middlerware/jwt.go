package middlerware

import (
	"assistantor/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, err := auth.VerifyToken(ctx); err == nil {
			ctx.Next()
		} else {
			ctx.JSON(http.StatusOK, gin.H{"code": 4001})
			ctx.Abort()
		}
	}
}
