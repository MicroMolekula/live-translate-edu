package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/live-translate-edu/internal/dto"
	"github.com/live-translate-edu/internal/services"
	"github.com/live-translate-edu/internal/utils"
	"net/http"
)

type AuthController struct {
	authService *services.AuthService
	roomService *services.RoomService
}

func NewAuthController(service *services.AuthService, room *services.RoomService) *AuthController {
	return &AuthController{
		authService: service,
		roomService: room,
	}
}

// Auth Авторизация
//
// @Tags user
// @Summary Авторизация
// @Description Авторизация в систему
// @Param credentials body dto.AuthDto true "почта и пароль"
// @Success 200 {object} TokenResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth [post]
func (ac *AuthController) Auth(ctx *gin.Context) {
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

	ctx.JSON(http.StatusOK, TokenResponse{
		Token:   token,
		Success: true,
	})
}

// Me Получение данных о авторизованном пользователе
//
// @Tags user
// @Summary Получение данных о авторизованном пользователе
// @Description Получение данных о авторизованном пользователе по jwt
// @Security ApiKeyAuth
// @Success 200 {object} dto.UserDTO
// @Failure 500 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /me [get]
func (ac *AuthController) Me(ctx *gin.Context) {
	user, ok := ctx.Value("user").(*dto.UserDTO)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, errors.New("internal server error"), "Error server")
	}
	ctx.JSON(http.StatusOK, user)
}

func (ac *AuthController) AuthRequiredMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := utils.GetTokenFromHeader(ctx)
		if err != nil {
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
		ctx.Set("jwt", tokenString)
		ctx.Next()
	}
}

// Users Получение данных о всех пользователях
//
// @Tags admin
// @Summary Получение данных о всех пользователях
// @Description Получение данных о всех пользователях доступно (доступно только админу)
// @Security ApiKeyAuth
// @Success 200 {array} dto.UserDTO
// @Failure 500 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /admin/users [get]
func (ac *AuthController) Users(ctx *gin.Context) {
	users, err := ac.authService.AllUsers()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err, "Internal Server Error")
		return
	}
	ctx.JSON(http.StatusOK, users)
}
