CREATE TABLE books(
    book_id SERIAL, 
    title VARCHAR(250) NOT NULL, 
    author VARCHAR(100) NOT NULL, 
    publisher VARCHAR(100) NOT NULL, 
    summary VARCHAR(250) NOT NULL, 
    stock INT NOT NULL, 
    max_stock INT NOT NULL, 
    created_at TIMESTAMP NOT NULL, 
    updated_at TIMESTAMP NOT NULL, 
    PRIMARY KEY(book_id));