CREATE TABLE items (
	id SERIAL PRIMARY KEY NOT NULL,
	name TEXT NOT NULL,
	hash_name TEXT NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL
);

CREATE TABLE alerts_daily_schedule (
	id SERIAL PRIMARY KEY NOT NULL,
	item_id INTEGER NULL,
	is_active bool DEFAULT false NULL,
	created_at timestamptz DEFAULT now() NULL,
	discord_id BIGINT NULL
);
