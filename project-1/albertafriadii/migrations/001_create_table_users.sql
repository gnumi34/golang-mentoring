CREATE TABLE users(
    user_id SERIAL NOT NULL, 
    username VARCHAR(50) NOT NULL, 
    email VARCHAR(100) NOT NULL, 
    password VARCHAR(100) NOT NULL, 
    created_at TIMESTAMP NOT NULL, 
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY(user_id));