package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/live-translate-edu/internal/dto"
	"net/http"
	"slices"
)

func RoleMiddleware(roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, ok := ctx.Value("user").(*dto.UserDTO)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}
		if !slices.Contains(roles, user.Role) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "forbidden",
				"message": "Доступ запрещен",
			})
			return
		}
		ctx.Next()
	}
}
