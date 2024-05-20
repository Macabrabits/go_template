package db

import (
	"database/sql"
	"fmt"
	"github.com/macabrabits/go_template/configs"
)

func Initialize() (*sql.DB, error) {
	cfg := configs.GetConfig().MysqlCFG
	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, err
	}
	fmt.Println("Database Connected!")
	return db, nil
}
