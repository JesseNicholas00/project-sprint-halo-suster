CREATE TABLE users (
    user_id TEXT PRIMARY KEY,
    nip BIGINT UNIQUE,
    is_admin BOOLEAN,
    name TEXT,
    password TEXT,
    active BOOLEAN,
    image_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
