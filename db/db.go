package db

import (
	"database/sql"

	"echo_rest/config"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()
	connStr := "host=" + conf.DB_HOST + " port=" + conf.DB_PORT + " user=" + conf.DB_USERNAME + " " + "password=" + conf.DB_PASSWORD + " dbname=" + conf.DB_NAME + " sslmode=disable"
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err.Error() + " connection string error")
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error() + " DSN error")
	}
}

func CreatCon() *sql.DB {
	return db
}
