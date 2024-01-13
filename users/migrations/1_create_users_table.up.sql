CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    roles VARCHAR(255) DEFAULT 'basic',
    active BOOLEAN DEFAULT true,
    verified BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_id UNIQUE (id),
    CONSTRAINT unique_email UNIQUE (email)
);
-- Index on email column
CREATE INDEX IF NOT EXISTS idx_email ON users(email);