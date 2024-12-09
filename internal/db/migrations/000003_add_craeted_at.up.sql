
ALTER TABLE sessions ADD COLUMN created_at timestamptz NOT NULL DEFAULT now();