package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/macabrabits/go_template/configs"
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
	fmt.Println("Connectedd!")
	return db
}
