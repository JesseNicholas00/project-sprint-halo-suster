CREATE TABLE patients (
    identity_number BIGINT PRIMARY KEY,
    phone_number TEXT,
    name TEXT,
    birth_date TIMESTAMP,
    gender TEXT,
    image_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
