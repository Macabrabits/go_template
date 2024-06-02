package repository

import (
	"context"
	"database/sql"

	"github.com/macabrabits/go_template/db/sqlc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// type CatRepository interface {
// 	list() []sqlc.Cat
// 	get(id string) sqlc.Cat
// 	create() sqlc.Cat
// 	update(id string, cat sqlc.Cat) sqlc.Cat
// 	remove(id string)
// }

type CatRepository struct {
	db *sql.DB
}

var (
	tracer = otel.Tracer("CatsRepository")
)

func NewCatRepository(db *sql.DB) CatRepository {
	return CatRepository{db}
}

func (c CatRepository) Create(ctx context.Context, params sqlc.CreateCatParams) (sql.Result, error) {
	_, span := tracer.Start(ctx, "DB CreateCat", trace.WithAttributes(attribute.String("db", "mysql")))
	defer span.End()
	queries := sqlc.New(c.db)
	result, err := queries.CreateCat(ctx, params)
	return result, err
}

func (c CatRepository) List(ctx context.Context) ([]sqlc.Cat, error) {
	_, span := tracer.Start(ctx, "DB ListCats", trace.WithAttributes(attribute.String("db", "mysql")))
	defer span.End()
	queries := sqlc.New(c.db)
	cats, err := queries.ListCats(ctx)
	return cats, err
}

func (c CatRepository) get(ctx context.Context, id uint64) (sqlc.Cat, error) {
	queries := sqlc.New(c.db)
	cat, err := queries.GetCat(ctx, id)
	return cat, err
}
