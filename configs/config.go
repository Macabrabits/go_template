package configs

import (
	"github.com/go-sql-driver/mysql"
	"os"
)

type Config struct {
	Port     string
	MysqlCFG mysql.Config
}

func config(defaultValue string, optionalValue string) string {
	if optionalValue == "" {
		return defaultValue
	}
	return optionalValue
}

func GetConfig() Config {
	cfg := Config{
		Port: config("8080", os.Getenv("PORT")),
		MysqlCFG: mysql.Config{
			User:   config("root", os.Getenv("DBUSER")),
			Passwd: config("root", os.Getenv("DBPASS")),
			Net:    "tcp",
			Addr:   config("localhost", os.Getenv("DBHOST")) + ":3306",
			DBName: "app",
		},
	}
	return cfg
}
