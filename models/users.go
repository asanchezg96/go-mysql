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
	Id       int64
	Name     string
	Password string
	Email    string
}

// List data
type Users []User

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
	result, _ := database.Exec(sql, user.Name, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

// List data from table users
func ListUsers() Users {
	sql := "SELECT id, name, password, email FROM users"
	users := Users{}
	rows, _ := database.Query(sql)
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users
}
