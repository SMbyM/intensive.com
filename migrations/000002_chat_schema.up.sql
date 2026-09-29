CREATE TABLE chat (
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE messages (
    id         SERIAL PRIMARY KEY,
    ch         INT NOT NULL REFERENCES chat(id) ON DELETE CASCADE,
    fr         INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    message    TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);