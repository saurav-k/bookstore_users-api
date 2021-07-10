package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	dbUsername = "db_username"
	dbPassword = "db_password"
	dbHost = "db_host"
	dbSchema = "db_schema"
)

var (
	Client *sql.DB
	username = os.Getenv(dbUsername)
	password = os.Getenv(dbPassword)
	host = os.Getenv(dbHost)
	schema = os.Getenv(dbSchema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root", "Fanatics@123", "127.0.0.1:3306", "users_db",
	)
	log.Println(dataSourceName)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database successfully configured")
	
}