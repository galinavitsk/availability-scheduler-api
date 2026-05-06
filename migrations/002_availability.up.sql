CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS availability (
    id             UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    slug           TEXT        NOT NULL,
    name           TEXT        NOT NULL,
    local_timezone TEXT        NOT NULL,
    slots_by_date  JSONB       NOT NULL DEFAULT '{}',
    hero_class     TEXT        NOT NULL,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_availability_slug ON availability(slug);
