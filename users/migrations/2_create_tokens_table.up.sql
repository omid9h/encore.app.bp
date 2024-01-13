CREATE TABLE tokens (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token TEXT NOT NULL,
    CONSTRAINT unique_user_id UNIQUE (user_id)
);
-- Index on user_id column
CREATE INDEX IF NOT EXISTS idx_user_id ON tokens(user_id);