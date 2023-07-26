package main

import (
	"fmt"
	"go-mysql-crud/database"
	"go-mysql-crud/models"
)

func main() {
	fmt.Println("GO-MYSQL-CRUD")
	database.Connect()
	database.Ping()

	database.CreateTable(models.Schema)

	//Verify if table exists
	//database.ExistsTable("users")
	database.Close()
}
