CREATE DATABASE IF NOT EXISTS name_db;
USE name_db;
CREATE USER IF NOT EXISTS 'user'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON name_db.* TO 'name_db'@'localhost';
FLUSH PRIVILEGES;
