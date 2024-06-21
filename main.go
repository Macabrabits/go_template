package main

import (
	"context"
	// "errors"
	"os"
	"os/signal"

	"github.com/coreos/go-oidc"
	"github.com/macabrabits/go_template/controller"
	"github.com/macabrabits/go_template/db"
	"github.com/macabrabits/go_template/repository"
	"github.com/macabrabits/go_template/router"
	"github.com/macabrabits/go_template/services"
)

var (
	keycloakIssuerURL = "http://mykeycloak:8080/realms/app"
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

//	@host		localhost:8082
//	@BasePath	/api/v1

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

//	@securitydefinitions.oauth2.implicit	OAuth2Implicit
//	@tokenUrl								http://mykeycloak:8080/realms/app/protocol/openid-connect/token
//	@authorizationurl						http://mykeycloak:8080/realms/app/protocol/openid-connect/auth
//	@scope.openid							Grants read access

// //	@securityDefinitions.oauth2.application	OAuth2Application
// //	@tokenUrl								http://mykeycloak:8080/realms/app/protocol/openid-connect/token
// //	@authorizationurl						http://mykeycloak:8080/realms/app/protocol/openid-connect/auth
// //	@scope.openid							Grants read access

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// // Set up OpenTelemetry.
	// otelShutdown, err := setupOTelSDK(ctx)
	// if err != nil {
	// 	return
	// }
	// // Handle shutdown properly so nothing leaks.
	// defer func() {
	// 	err = errors.Join(err, otelShutdown(context.Background()))
	// }()

	db, err := db.Initialize()
	if err != nil {
		panic(err)
	}

	// Create OIDC provider
	oidcProvider, err := oidc.NewProvider(ctx, keycloakIssuerURL)
	if err != nil {
		panic(err)
	}

	catsRepository := repository.NewCatRepository(db)
	catsService := services.NewCatsService(&catsRepository)
	catsController := controller.NewCatsController(&catsService)
	authService := services.NewAuthsService()
	authController := controller.NewAuthController(&authService)
	auth2Controller := controller.NewAuth2Controller(oidcProvider)

	router.Initialize(&catsController, &authController, &auth2Controller)
}
