CREATE UNIQUE INDEX users_nip ON users(nip);
CREATE INDEX users_role ON users(role);
CREATE INDEX users_created_at ON users(created_at);
