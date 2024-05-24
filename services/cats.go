package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/macabrabits/go_template/db/sqlc"

	// "go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

type CatsService struct {
	db *sql.DB
}

type Cat struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"  validate:"required"`
	Age   int8   `json:"age"   validate:"required,gte=0,lte=25"`
	Breed string `json:"breed" validate:"required"`
}

// var sqlcNew = sqlc.New
const name = "Cats"

var (
	// tracer = otel.Tracer(name)
	meter = otel.Meter(name)
	// logger =
	rollCnt metric.Int64Counter
)

func NewCatsService(db *sql.DB) CatsService {
	return CatsService{db}
}

func (svc *CatsService) GetCats() (gin.H, error) {
	fmt.Println("midtrace")
	ctx := context.Background()
	queries := sqlc.New(svc.db)
	cats, err := queries.ListCats(ctx)
	if err != nil {
		return gin.H{}, err
	}

	return gin.H{
		"message": "success",
		"data":    cats,
	}, err
}

func (svc *CatsService) CreateCat(cat sqlc.CreateCatParams) (gin.H, error) {
	ctx := context.Background()
	queries := sqlc.New(svc.db)
	//Insert in the DB
	result, err := queries.CreateCat(ctx, cat)
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
