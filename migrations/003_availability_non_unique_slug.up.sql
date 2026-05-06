ALTER TABLE availability DROP CONSTRAINT IF EXISTS availability_slug_key;

DROP INDEX IF EXISTS idx_availability_slug;

CREATE INDEX IF NOT EXISTS idx_availability_slug ON availability(slug);
