package services

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/macabrabits/go_template/db/sqlc"
	"github.com/macabrabits/go_template/repository"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

type CatsService struct {
	// db         *sql.DB
	repository *repository.CatRepository
}

type Cat struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"  validate:"required"`
	Age   int8   `json:"age"   validate:"required,gte=0,lte=25"`
	Breed string `json:"breed" validate:"required"`
}

const name = "Cats"

var (
	// tracer = otel.Tracer(name)
	meter = otel.Meter(name)
	// logger =
	rollCnt metric.Int64Counter
)

func NewCatsService(
	repository *repository.CatRepository,
) CatsService {
	return CatsService{
		repository,
	}
}

func (svc *CatsService) GetCats(ctx context.Context) (gin.H, error) {
	cats, err := svc.repository.List(ctx)
	if err != nil {
		return nil, err
	}

	return gin.H{
		"message": "success",
		"data":    cats,
	}, nil
}

func (svc *CatsService) CreateCat(ctx context.Context, params sqlc.CreateCatParams) (gin.H, error) {
	//Insert in the DB
	result, err := svc.repository.Create(ctx, params)
	if err != nil {
		return nil, err
	}
	//Add metric
	rollCnt, err = meter.Int64Counter("created_cats", metric.WithDescription("Created cats"))
	if err != nil {
		return nil, fmt.Errorf("error saving the metric: %w", err)
	}
	rollCnt.Add(context.Background(), 1)

	id, err := result.LastInsertId()
	res := gin.H{
		"message": "cat create successfully",
		"data":    gin.H{"id": id},
	}
	return res, err
}
