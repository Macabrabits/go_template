package main

import (
	"context"
	"errors"

	// "log"
	// "net"
	// "net/http"
	"os"
	"os/signal"

	// "time"

	"github.com/macabrabits/go_template/controller"
	"github.com/macabrabits/go_template/db"
	"github.com/macabrabits/go_template/router"
	"github.com/macabrabits/go_template/services"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
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
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	otelShutdown, err := setupOTelSDK(ctx)
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()
	meter := otel.Meter("test-meter")
	meter.Int64Counter("run", metric.WithDescription("The number of times the iteration ran"))

	db, err := db.Initialize()
	if err != nil {
		panic(err)
	}

	catsService := services.NewCatsService(db)
	catsController := controller.NewCatsController(&catsService)
	router.Initialize(&catsController)

}
