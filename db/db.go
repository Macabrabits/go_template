package db

import (
	"database/sql"
	"fmt"
	"github.com/macabrabits/go_template/configs"
	"log"
)

func Db() *sql.DB {
	cfg := configs.GetConfig().MysqlCFG
	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db
}
