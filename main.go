package main

import (
	"fmt"
	"go-mysql-crud/database"
)

func main() {
	fmt.Println("GO-MYSQL-CRUD")
	database.Connect()
	database.Ping()

	//database.CreateTable(models.Schema)

	//Verify if table exists
	database.ExistsTable("users")
	database.Close()
}
