package main

import (
	"github.com/macabrabits/go_template/controller"
	"github.com/macabrabits/go_template/db"
	"github.com/macabrabits/go_template/router"
	"github.com/macabrabits/go_template/services"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	db, err := db.Initialize()
	if err != nil {
		panic(err)
	}
	catsService := services.NewCatsService(db)
	catsController := controller.NewCatsController(&catsService)

	router.Initialize(&catsController)
}
