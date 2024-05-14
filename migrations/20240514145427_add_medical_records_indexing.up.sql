CREATE INDEX medical_records_identity_number ON medical_records(patient_identity_number);
CREATE INDEX medical_records_created_by_nip ON medical_records(created_by_nip);
CREATE INDEX medical_records_created_at ON medical_records(created_at);
