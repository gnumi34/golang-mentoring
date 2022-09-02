CREATE TABLE public.book_collection (
    book_id SERIAL PRIMARY KEY NOT NULL,
    title VARCHAR(50) NOT NULL,
    author VARCHAR(50) NOT NULL,
    publisher VARCHAR(50),
    book_summary VARCHAR(150),
    book_stock INT,
    max_book_stock INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE public.book_lending (
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

CREATE TABLE public.book_request_approval (
    approval_id SERIAL PRIMARY KEY NOT NULL,
    book_id INT,
    user_id INT,
    request_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_accepted BOOLEAN,
    notes VARCHAR(100),
    FOREIGN KEY (book_id) REFERENCES book_collection(book_id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);