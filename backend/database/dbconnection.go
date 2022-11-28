package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Dbconnection() (*sql.DB, error) {
	var dbConfig map[string]string
	dbConfig, failed := godotenv.Read()

	if failed != nil {
		return nil, failed
	}

	urlConnection := dbConfig["MYSQL_USER"] + ":" + dbConfig["MYSQL_PASSWORD"] + "@" + dbConfig["MYSQL_PROTOCOL"] + "(" + dbConfig["MYSQL_HOST"] + ":" + dbConfig["MYSQL_PORT"] + ")/" + dbConfig["MYSQL_DBNAME"]

	db, failed := sql.Open("mysql", urlConnection)

	if failed != nil {
		return nil, failed
	}

	return db, nil
}
