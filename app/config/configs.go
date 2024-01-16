package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const driverName = "mysql"
const accessTokenTemplate = "%s:%s@tcp(%s:%s)/%s?parseTime=true"

var Conn *sql.DB

func CreateDBConnection() (*sql.DB, error) {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DATABASE")

	var err error
	Conn, err = sql.Open(driverName, fmt.Sprintf(accessTokenTemplate, user, password, host, port, database))
	if err != nil {
		return nil, err
	}

	if err := Conn.Ping(); err != nil {
		return nil, fmt.Errorf("can't connect to mysql server."+
			"MYSQL_USER=%s, "+
			"MYSQL_PASSWORD=%s, "+
			"MYSQL_HOST=%s, "+
			"MYSQL_PORT=%s, "+
			"MYSQL_DATABASE=%s, "+
			"error=%+v",
			user, password, host, port, database, err)
	}

	return Conn, nil
}
