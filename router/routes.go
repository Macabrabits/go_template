package router

import (
	// docs "github.com/arthur404dev/gopportunities/docs"
	"github.com/gin-gonic/gin"
	"github.com/macabrabits/go_template/controller"
	docs "github.com/macabrabits/go_template/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath + "/cats")
	{
		v1.GET("/", controller.GetCats)
		v1.POST("/", controller.CreateCat)
		// v1.DELETE("/opening", handler.DeleteOpeningHandler)
		// v1.PUT("/opening", handler.UpdateOpeningHandler)
		// v1.GET("/openings", handler.ListOpeningsHandler)

	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
