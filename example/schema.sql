DROP TABLE IF EXISTS "users" CASCADE;
DROP TABLE IF EXISTS "vilages" CASCADE;
DROP TABLE IF EXISTS "people" CASCADE;
DROP TABLE IF EXISTS "cities" CASCADE;
DROP TABLE IF EXISTS "countries" CASCADE;


-- TABLE "users" --
CREATE TABLE IF NOT EXISTS users (
  "email"     CITEXT UNIQUE NOT NULL PRIMARY KEY,
  "firstname" CITEXT NOT NULL,
  "lastname"  CITEXT NOT NULL,
  "phone"     CITEXT NOT NULL,
  "about"     TEXT,
  "birthday"  timestamp not null
);

-- TABLE "vilages" --
CREATE TABLE IF NOT EXISTS vilages (
  "name"       CITEXT NOT NULL,
  "population" INTEGER DEFAULT 0
);

-- TABLE "people" --
create table people (
  "id"        serial not null primary key,
  "name"      varchar(256) not null,
  "lastname"  varchar(256) not null,
  "birthday"  date not null,
  "some_flag" integer not null,
  "created"   timestamp(0) not null default current_timestamp
);

-- TABLE "cities" --
create table cities (
  "id"         serial not null primary key,
  "name"       varchar(256) not null,
  "country_id" integer not null,
  "created"    timestamp(0) not null default current_timestamp
);

-- TABLE "countries" --
create table countries (
  "id"      serial not null primary key,
  "name"    varchar(256) not null,
  "created" timestamp(0) not null default current_timestamp
);
