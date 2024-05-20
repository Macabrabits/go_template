package services

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	sqlc "github.com/macabrabits/go_template/sqlc"
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

func NewCatsService(db *sql.DB) CatsService {
	return CatsService{db}
}

func (s *CatsService) GetCats() (gin.H, error) {
	ctx := context.Background()
	queries := sqlc.New(s.db)
	cats, err := queries.ListCats(ctx)
	if err != nil {
		return gin.H{}, err
	}

	return gin.H{
		"message": "success",
		"data":    cats,
	}, err
}

func (s *CatsService) CreateCat(cat sqlc.CreateCatParams) (gin.H, error) {
	ctx := context.Background()
	queries := sqlc.New(s.db)
	result, err := queries.CreateCat(ctx, cat)
	if err != nil {
		return gin.H{}, err
	}
	id, err := result.LastInsertId()
	res := gin.H{
		"message": "cat create successfully",
		"data":    gin.H{"id": id},
	}
	return res, err
}
