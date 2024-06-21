package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/macabrabits/go_template/db/sqlc"
	"github.com/macabrabits/go_template/services"
	"go.opentelemetry.io/otel"
)

type CatsController struct {
	svc services.CatsService
}
type JSONResultList struct {
	Message string         `json:"message"`
	Data    []services.Cat `json:"data"`
}
type JSONResult struct {
	Message string       `json:"message"`
	Data    services.Cat `json:"data"`
}

func NewCatsController(s *services.CatsService) CatsController {
	return CatsController{*s}
}

var (
	tracer = otel.Tracer("testapp")
)

// GetCats godoc
//
//	@Summary		Get all cats
//	@Description	Get all cats
//	@Tags			Cats
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	JSONResultList
//	@Router			/cats [get]
//	@Security		OAuth2Implicit
func (s *CatsController) GetCats(ctx *gin.Context) {
	spanName := ctx.Request.Method + " - " + ctx.Request.URL.Path
	tctx, span := tracer.Start(ctx, spanName)
	defer span.End()

	res, err := s.svc.GetCats(tctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// CreateCat godoc
//
//	@Summary		Create Cat
//	@Description	Insert a Cat
//	@Tags			Cats
//	@Param			request	body	services.Cat	true	"cat"	services.Cat
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	JSONResult
//	@Router			/cats [post]
//	@Security		OAuth2Implicit
func (s *CatsController) CreateCat(ctx *gin.Context) {
	spanName := ctx.Request.Method + " - " + ctx.Request.URL.Path
	tctx, span := tracer.Start(ctx, spanName)
	defer span.End()

	var cat sqlc.CreateCatParams

	if err := ctx.Bind(&cat); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err := validator.New().Struct(cat)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}

	res, err := s.svc.CreateCat(tctx, cat)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
