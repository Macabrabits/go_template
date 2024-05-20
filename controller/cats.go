package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/macabrabits/go_template/db/sqlc"
	"github.com/macabrabits/go_template/services"
	"net/http"
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

// GetCats godoc
//
//	@Summary		Get all cats
//	@Description	Get all cats
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	JSONResultList
//	@Router			/cats [get]
func (s *CatsController) GetCats(ctx *gin.Context) {
	res, err := s.svc.GetCats()
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
//	@Tags			accounts
//	@Param			request	body	services.Cat	true	"cat"	services.Cat
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	JSONResult
//	@Router			/cats [post]
func (s *CatsController) CreateCat(ctx *gin.Context) {
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

	res, err := s.svc.CreateCat(cat)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
