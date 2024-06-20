package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/macabrabits/go_template/services"
	"net/http"
)

type AuthController struct {
	svc services.AuthService
}
type AuthResponse struct {
	Message string `json:message`
	Data    string `json:data`
}

func NewAuthController(s *services.AuthService) AuthController {
	return AuthController{*s}
}

// Auth godoc
//
//	@Summary		GetToken
//	@Description	GetToken
//	@Tags			Auth
//	@Param			request	body	services.AuthParams	true	"AuthParams"	services.AuthParams
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	services.AuthResponse
//	@Router			/auth/gettoken [post]
func (s *AuthController) GetToken(ctx *gin.Context) {
	spanName := ctx.Request.Method + " - " + ctx.Request.URL.Path
	tctx, span := tracer.Start(ctx, spanName)
	defer span.End()
	authParams := services.AuthParams{}
	if err := ctx.Bind(&authParams); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := s.svc.GetToken(tctx, authParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// AuthCallback godoc
//
//	@Summary		AuthCallback
//	@Description	AuthCallback
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	AuthResponse
//	@Router			/auth/callback [post]
func (s *AuthController) AuthCallback(ctx *gin.Context) {
	spanName := ctx.Request.Method + " - " + ctx.Request.URL.Path
	tctx, span := tracer.Start(ctx, spanName)
	defer span.End()

	fmt.Println(ctx.Request.Method)
	fmt.Println(ctx.Request.Body)
	fmt.Println(ctx.Request.URL)

	res, err := s.svc.AuthCallback(tctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// Auth godoc
//
//	@Summary		Auth
//	@Description	Auth
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	AuthResponse
//	@Router			/auth [get]
//	@Security		BearerAuth
func (s *AuthController) Auth(ctx *gin.Context) {
	spanName := ctx.Request.Method + " - " + ctx.Request.URL.Path
	tctx, span := tracer.Start(ctx, spanName)
	defer span.End()

	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		ctx.Abort()
		return
	}

	_, err := s.svc.Auth(tctx, ctx.GetHeader("Authorization"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}

	// ctx.JSON(http.StatusOK, res)
	ctx.Next()
}
