CREATE DATABASE goweb_db;

CREATE TABLE users(
    id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(30) NOT NULL,
    password VARCHAR(64) NOT NULL,
    email VARCHAR(50),
    create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

show tables like 'users';
