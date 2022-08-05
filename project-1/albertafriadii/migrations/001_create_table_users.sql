CREATE TABLE users(
    user_id uuid DEFAULT uuid_generate_v4(), 
    username VARCHAR(50) NOT NULL, 
    email VARCHAR(100) NOT NULL, 
    password VARCHAR(100) NOT NULL, 
    created_at TIMESTAMP NOT NULL, 
    updated_at TIMESTAMP NOT NULL, 
    PRIMARY KEY(user_id));