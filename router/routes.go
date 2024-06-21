package router

import (
	"github.com/gin-gonic/gin"
	"github.com/macabrabits/go_template/controller"
	docs "github.com/macabrabits/go_template/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(
	router *gin.Engine,
	catsController *controller.CatsController,
	authController *controller.AuthController,
	auth2Controller *controller.Auth2Controller,

) {
	// handler.InitializeHandler()

	// router.Use(instrument)

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	cats := router.Group(basePath + "/cats")
	cats.Use(auth2Controller.AuthMiddleware())
	// cats.Use(auth2Controller.RefreshTokenIfNeeded())
	// cats.Use(authController.Auth)
	{
		cats.GET("/", catsController.GetCats)
		cats.POST("/", catsController.CreateCat)
		// v1.DELETE("/opening", handler.DeleteOpeningHandler)
		// v1.PUT("/opening", handler.UpdateOpeningHandler)
		// v1.GET("/openings", handler.ListOpeningsHandler)

	}
	auth := router.Group(basePath + "/auth")
	{
		auth.POST("/gettoken", authController.GetToken)
		auth.GET("/callback", authController.AuthCallback)
		auth.GET("/", authController.Auth)
	}
	auth2 := router.Group(basePath + "/auth2")
	{

		auth2.GET("/", auth2Controller.IndexHandler)
		auth2.POST("/login", auth2Controller.LoginHandler())
		auth2.GET("/login", auth2Controller.LoginHandler())
		auth2.GET("/callback", auth2Controller.CallbackHandler())
		auth2.GET("/logout", auth2Controller.LogoutHandler)
	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
