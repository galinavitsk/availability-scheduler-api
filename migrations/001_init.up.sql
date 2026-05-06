CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS sessions (
    id             UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    name           TEXT        NOT NULL,
    start_time     TEXT        NOT NULL,
    end_time       TEXT        NOT NULL,
    slug           TEXT        NOT NULL UNIQUE,
    time_zone      TEXT        NOT NULL,
    selected_dates TEXT[]      NOT NULL DEFAULT '{}',
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_sessions_start_time ON sessions(start_time);
