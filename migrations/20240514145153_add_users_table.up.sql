CREATE TABLE users (
    nip BIGINT PRIMARY KEY,
    name TEXT,
    password TEXT,
    role TEXT,
    active BOOLEAN,
    image_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
