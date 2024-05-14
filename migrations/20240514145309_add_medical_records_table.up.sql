CREATE TABLE medical_records (
    patient_identity_number BIGINT PRIMARY KEY,
    patient_phone_number TEXT,
    patient_name TEXT,
    patient_birth_date TEXT,
    patient_gender TEXT,
    symptoms TEXT,
    medication TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by_user_id TEXT,
);
