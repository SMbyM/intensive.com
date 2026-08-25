CREATE TABLE users (
    id            SERIAL PRIMARY KEY,
    email         TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    name          TEXT NOT NULL,
    lastname      TEXT NOT NULL,
    nickname      TEXT,
    phone         TEXT,
    male          TEXT,
    birthday      DATE
);

CREATE TABLE friends (
    fst INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    snd INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (fst, snd)
);