CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR (50) NOT NULL,
    avatar_url TEXT NOT NULL,
    discord_uid BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS games (
    id BIGSERIAL PRIMARY KEY,
    game_name VARCHAR (50) NOT NULL,
    game_icon_url VARCHAR (200) NOT NULL
);

CREATE TABLE IF NOT EXISTS sessions (
    id BIGSERIAL PRIMARY KEY,
    is_finish BOOLEAN DEFAULT false,
    session_end TIMESTAMP,
    session_start TIMESTAMP,
    created_by BIGINT NOT NULL REFERENCES users(id),
    game_id BIGINT NOT NULL REFERENCES games(id)
);


CREATE TABLE IF NOT EXISTS users_session (
    user_id BIGINT NOT NULL REFERENCES users(id),
    session_id BIGINT NOT NULL REFERENCES sessions(id)
);