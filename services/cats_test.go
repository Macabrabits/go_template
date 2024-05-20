package services

import (
	// "context"
	"context"
	"database/sql"
	"testing"

	"github.com/macabrabits/go_template/db/sqlc"
	// "github.com/macabrabits/go_template/sqlc"
)

func TestTest(t *testing.T) {
	param := "world of tests"
	// msg := toTest(param)
	msg := "todo"
	expected := "hello world of tests"
	if msg != expected {
		t.Errorf("toTest(%q) = %q; want %q", param, msg, expected)
	}
}

type ResultMock interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}
type driverResult struct {
	resi ResultMock
}

func (dr driverResult) LastInsertId() (int64, error) {
	return 1, nil
}
func (dr driverResult) RowsAffected() (int64, error) {
	return 10, nil
}

type QueriesMock struct {
	db sqlc.DBTX
}

func (q QueriesMock) CreateCat(ctx context.Context, arg sqlc.CreateCatParams) (sql.Result, error)
func (q QueriesMock) ListCats(ctx context.Context) ([]sqlc.Cat, error) {
	return []sqlc.Cat{
		sqlc.Cat{ID: 2, Name: "Fa", Age: 22, Breed: "F"},
	}, nil
}
func (q QueriesMock) WithTx(tx *sql.Tx) *sqlc.Queries

// func (q *QueriesMock) ListCats() {
// 	return
// }

func TestGetCats(t *testing.T) {

	//mock
	// sqlcNew =
	// defer func() { sqlcNew = old }()
	// sqlcNew = func() {
	// 	return QueriesMock{}.db
	// }

	var db *sql.DB
	svc := NewCatsService(db)
	svc.GetCats()

}

// func TestCreateCat() {

// }
