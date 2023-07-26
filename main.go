package main

import (
	"fmt"
	"go-mysql-crud/database"
)

func main() {
	fmt.Println("GO-MYSQL-CRUD")
	database.Connect()
	database.Ping()
}
