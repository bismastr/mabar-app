ALTER TABLE users_session DROP COLUMN discord_uid;

ALTER TABLE users_session ALTER COLUMN user_id SET NOT NULL;