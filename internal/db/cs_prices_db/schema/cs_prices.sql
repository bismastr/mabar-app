CREATE TABLE items (
	id SERIAL PRIMARY KEY NOT NULL,
	name TEXT NOT NULL,
	hash_name TEXT NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL
);