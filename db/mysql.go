package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func CreateInstance(config interface{}, dbname string) *sql.DB {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			config.(map[string]interface{})["username"],
			config.(map[string]interface{})["password"],
			config.(map[string]interface{})["hostname"],
			config.(map[string]interface{})["port"],
			dbname,
		),
	)

	if err != nil {
		panic(err)
	}

	return db
}
