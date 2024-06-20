package router

import (
	"github.com/gin-gonic/gin"
	"github.com/macabrabits/go_template/configs"
	"github.com/macabrabits/go_template/controller"
)

func Initialize(
	catsCotnroller *controller.CatsController,
	authCotnroller *controller.AuthController,
	auth2Cotnroller *controller.Auth2Controller,
) {
	// Initialize Router
	router := gin.Default()

	err := router.SetTrustedProxies(nil) //TODO: understand the utility of that
	if err != nil {
		panic(err)
	}

	cfg := configs.GetConfig()

	// Initialize Routes
	initializeRoutes(
		router,
		catsCotnroller,
		authCotnroller,
		auth2Cotnroller,
	)

	// Run the server
	router.Run("0.0.0.0:" + cfg.Port)
}
