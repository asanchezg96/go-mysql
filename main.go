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

	//Truncate table
	database.TruncateTable("users")

	//create new user
	//user := models.NewUser("Jhon", "123456", "jhon@gmai.com")
	user := models.CreateUser("Tony", "pass", "tony@gmail.com")
	fmt.Println(user)
	user2 := models.CreateUser("Jhon", "pass", "jhon@gmail.com")
	fmt.Println(user2)
	fmt.Println(models.ListUsers())
	user2.Name = "Jhonatan"
	user2.Save()
	fmt.Println(models.ListUsers())

	database.Close()
}
