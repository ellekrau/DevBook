CREATE DATABASE IF NOT EXISTS devbook;

USE mysql;

CREATE USER IF NOT EXISTS 'devbook' IDENTIFIED BY 'devbook';
GRANT ALL PRIVILEGES ON devbook.* TO 'devbook';

USE devbook;

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INT auto_increment PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    nickname VARCHAR(20) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    createdAt TIMESTAMP DEFAULT current_timestamp()
) ENGINE=INNODB;

DESC users;