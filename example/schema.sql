DROP TABLE IF EXISTS "people" CASCADE;
DROP TABLE IF EXISTS "cities" CASCADE;

-- TABLE "people" --
CREATE TABLE IF NOT EXISTS people (
  "nickname"  CITEXT UNIQUE PRIMARY KEY,
  "firstname" CITEXT UNIQUE NOT NULL,
  "lasname"   CITEXT UNIQUE NOT NULL,
  "email"     CITEXT UNIQUE NOT NULL,
  "phone"     CITEXT NOT NULL,
  "about"     TEXT
);

-- TABLE "cities" --
CREATE TABLE IF NOT EXISTS cities (
  "name"       CITEXT UNIQUE NOT NULL,
  "population" INTEGER DEFAULT 0
);