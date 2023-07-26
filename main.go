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
	//database.CreateTable(models.Schema)
	//Verify if table exists
	//database.ExistsTable("users")

	//create new user
	user := models.NewUser("Jhon", "123456", "jhon@gmai.com")
	fmt.Println(user)
	database.Close()
}
