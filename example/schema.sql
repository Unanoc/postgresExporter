DROP TABLE IF EXISTS "people" CASCADE;
DROP TABLE IF EXISTS "cities" CASCADE;

-- TABLE "people" --
CREATE TABLE IF NOT EXISTS people (
  "email"     CITEXT UNIQUE NOT NULL PRIMARY KEY,
  "firstname" CITEXT NOT NULL,
  "lastname"  CITEXT NOT NULL,
  "phone"     CITEXT NOT NULL,
  "about"     TEXT,
  "birthday"  timestamp not null
);

-- TABLE "cities" --
CREATE TABLE IF NOT EXISTS cities (
  "name"       CITEXT NOT NULL,
  "population" INTEGER DEFAULT 0
);