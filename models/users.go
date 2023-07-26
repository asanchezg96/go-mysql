package models

import "go-mysql-crud/database"

const Schema = `
	CREATE TABLE IF NOT EXISTS users (
		id INT (6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(50) NOT NULL, 
		email VARCHAR(50) NOT NULL,
		password VARCHAR(50) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
`

//Insert data

// Create type date to use in insert en table users
type User struct {
	Id       int
	Name     string
	Password string
	Email    string
}

//Create constructor for nuew user

func NewUser(name, password, email string) *User {
	user := &User{
		Name:     name,
		Password: password,
		Email:    email,
	}
	return user
}

// Create function to create user in tabble users
func CreateUser(name, password, email string) *User {
	user := NewUser(name, password, email)
	user.insert()
	return user
}

// Create method to insert data in table users
func (user *User) insert() {
	sql := "INSERT INTO users SET name= ?, password=?, email=?"
	_, err := database.Exec(sql, user.Name, user.Password, user.Email)
	if err != nil {
		panic(err)
	}
}
