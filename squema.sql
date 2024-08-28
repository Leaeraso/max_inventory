CREATE DATABASE max_inventory;
USE max_inventory;

CREATE TABLE users (
  id INT NOT NULL AUTO_INCREMENT,
  email VARCHAR(255) NOT NULL UNIQUE,
  name VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE products (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  price FLOAT NOT NULL,
  created_by INT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE roles (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE user_role (
    id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  role_id INT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);

INSERT INTO roles (id, name) VALUES (1, 'admin');
INSERT INTO roles (id, name) VALUES (2, 'seller');
INSERT INTO roles (id, name) VALUES (3, 'customer');
