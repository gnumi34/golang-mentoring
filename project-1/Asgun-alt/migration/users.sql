CREATE TABLE public.users (
	id SERIAL NOT NULL PRIMARY KEY,
	username VARCHAR(50),
	email VARCHAR(50),
	password VARCHAR(100),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NULL,
	deleted_at TIMESTAMP NULL
)