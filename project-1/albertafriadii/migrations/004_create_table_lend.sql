CREATE TABLE lend_books(
    lend_id SERIAL,
    book_id INT,
    user_id INT,
    requested_at TIMESTAMP NOT NULL,
    status BOOLEAN,
    notes VARCHAR(250),
    PRIMARY KEY(lend_id),
    FOREIGN KEY(book_id) REFERENCES books(book_id),
    FOREIGN KEY(user_id) REFERENCES users(user_id)
);