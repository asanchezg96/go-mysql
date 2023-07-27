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

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	} else {
		fmt.Println("Database Close Success")
	}
}

func CreateTable(schema string) {
	if !ExistsTable("users") {
		if _, err := Exec(schema); err != nil {
			panic(err)
		} else {
			fmt.Println("Table created successfully")
		}
	} else {
		fmt.Println("Table already exists")
	}
}

func ExistsTable(nameTable string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", nameTable)
	rows, _ := Query(sql)
	if rows.Next() {
		fmt.Println("Table exists")
		return true
	} else {
		fmt.Println("Table not exists")
		return false
	}
}

// Polymorphism method Execute
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		panic(err)
	}
	return result, err
}

// Polymorphism method Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		panic(err)
	}
	return rows, err

}

//Truncate table

func TruncateTable(tableName string) {
	if ExistsTable(tableName) {
		sql := fmt.Sprintf("TRUNCATE TABLE %s", tableName)
		if _, err := Exec(sql); err != nil {
			panic(err)
		} else {
			fmt.Println("Table truncate successfully")
		}
	} else {
		fmt.Println("Table not exists")
	}

}

func DropTable(tableName string) {
	sql := fmt.Sprintf("DROP TABLE %s", tableName)
	if ExistsTable(tableName) {
		if _, err := Exec(sql); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Table deleted")
		}
	} else {
		fmt.Println("Table not exists")
	}

}
