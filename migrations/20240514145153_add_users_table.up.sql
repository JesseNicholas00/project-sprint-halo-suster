CREATE TABLE users (
    user_id TEXT PRIMARY KEY,
    nip BIGINT UNIQUE,
    name TEXT,
    password TEXT,
    role TEXT,
    active BOOLEAN,
    image_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
