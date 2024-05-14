CREATE TABLE medical_records (
    medical_record_id SERIAL PRIMARY KEY,
    patient_identity_number BIGINT,
    patient_phone_number TEXT,
    patient_name TEXT,
    patient_birth_date TEXT,
    patient_gender TEXT,
    symptoms TEXT,
    medication TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT
);
