package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/live-translate-edu/internal/configs"
	"github.com/live-translate-edu/internal/dto"
	"github.com/live-translate-edu/internal/services"
	"net/http"
	"strings"
)

type AuthController struct {
	jwtSecret   string
	authService *services.AuthService
}

func newAuthController() *AuthController {
	as := services.NewAuthService()
	return &AuthController{
		jwtSecret:   configs.Cfg.JWT.Secret,
		authService: as,
	}
}

func (ac *AuthController) auth(ctx *gin.Context) {
	var request dto.AuthDto

	if err := ctx.ShouldBindJSON(&request); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err, "Bad Request")
		return
	}

	token, err := ac.authService.Auth(request)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrorInvalidCredentials):
			newErrorResponse(ctx, http.StatusUnauthorized, err, "Invalid Credentials")
		default:
			newErrorResponse(ctx, http.StatusInternalServerError, err, "Internal Server Error")
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":   token,
		"success": true,
	})
}

func (ac *AuthController) me(ctx *gin.Context) {
	user, ok := ctx.Value("user").(*dto.UserDTO)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.New("internal server error"), "Error server")
	}
	ctx.JSON(http.StatusOK, user)
}

func (ac *AuthController) authRequiredMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		tokenString, ok := strings.CutPrefix(authHeader, "Bearer ")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "unauthorized",
				"message": "Invalid Auth Token",
			})
			return
		}
		user, err := ac.authService.VerifyToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   err.Error(),
				"message": "Invalid Auth Token",
			})
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}

func (ac *AuthController) users(ctx *gin.Context) {
	users, err := ac.authService.AllUsers()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	ctx.JSON(http.StatusOK, users)
}
