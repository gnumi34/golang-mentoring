CREATE TABLE borrow_books(
    borrow_id SERIAL,
    book_id INT,
    user_id INT,
    borrow_date TIMESTAMP NOT NULL,
    dead_line DATE,
    status BOOLEAN,
    return_date TIMESTAMP,
    notes VARCHAR(250),
    PRIMARY KEY(borrow_id),
    FOREIGN KEY(book_id) REFERENCES books(book_id),
    FOREIGN KEY(user_id) REFERENCES users(user_id)
);