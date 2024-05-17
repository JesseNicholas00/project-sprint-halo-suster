CREATE UNIQUE INDEX users_nip ON users(nip);
CREATE INDEX users_is_admin ON users(is_admin);
CREATE INDEX users_created_at ON users(created_at);
