package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const dataSourceName = "go_mysql_user:go_mysql_password@tcp(localhost:3306)/go_mysql_db"
const driverName = "mysql"

var db *sql.DB

func Connect() {

	conect, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		panic(err)
	}
	db = conect
	fmt.Println("Database Open  Success")
}

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("Database conection Success")
	}
}
