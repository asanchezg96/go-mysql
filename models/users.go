package models

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
