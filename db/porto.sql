DROP TABLE IF EXISTS ports;

CREATE TABLE ports (
  id        BIGSERIAL PRIMARY KEY,
  name      TEXT,
  city      TEXT,
	сountry   TEXT,
	alias     TEXT,
	regions   TEXT,
  lat       REAL,
  long      REAL,
  province  TEXT,
	timezone  TEXT,
	unlocks   TEXT,
	code      TEXT
);