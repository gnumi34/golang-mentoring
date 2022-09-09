CREATE TABLE public.book_collections (
    book_id SERIAL PRIMARY KEY NOT NULL,
    title VARCHAR(50) NOT NULL,
    author VARCHAR(50) NOT NULL,
    publisher VARCHAR(50),
    book_summary VARCHAR(500),
    book_stock INT,
    max_book_stock INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE public.borrowed_books (
    lend_id SERIAL PRIMARY KEY NOT NULL,
    book_id INT,
    user_id INT,
    lent_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    due_date DATE,
    status BOOLEAN,
    returned_at TIMESTAMP,
    notes VARCHAR(100),
    FOREIGN KEY (book_id) REFERENCES book_collection(book_id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE public.lend_books (
    approval_id SERIAL PRIMARY KEY NOT NULL,
    book_id INT,
    user_id INT,
    request_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_accepted BOOLEAN,
    notes VARCHAR(100),
    FOREIGN KEY (book_id) REFERENCES book_collection(book_id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

SELECT u.username AS User, bc.title AS Book_Title, COUNT(bc.title) AS Book_Total
    FROM borrowed_books bb
    JOIN users u ON u.id = bb.user_id
    JOIN book_collections bc ON bc.book_id = bb.book_id
GROUP BY u.username, bc.title