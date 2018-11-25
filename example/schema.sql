DROP TABLE IF EXISTS "people" CASCADE;
DROP TABLE IF EXISTS "cities" CASCADE;
DROP TABLE IF EXISTS "psqltypes" CASCADE;

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

-- TABLE "psqltypes" --
CREATE TABLE IF NOT EXISTS psqltypes (
  "bigint"        int8 NOT NULL,
  "bigserial"     serial8 NOT NULL,
  "bit"           bit[5] NOT NULL,
  "varbit"        varbit NOT NULL,
  "boolean"       boolean NOT NULL,
  "box"           box NOT NULL,
  "character"     char[5] NOT NULL,
  "cidr"          cidr NOT NULL,
  "circle"        circle NOT NULL,
  "date"          date NOT NULL,
  "float8"        float8 NOT NULL,
  "inet"          inet NOT NULL,
  "integer"       int NOT NULL,
  "interval"      interval[?] NOT NULL,
  "json"          json NOT NULL,
  "line"          line NOT NULL,
  "lseg"          lseg NOT NULL,
  "macaddr"       macaddr NOT NULL,
  "money"         int8 NOT NULL,
  "numeric"       decimal[?] NOT NULL,
  "path"          path NOT NULL,
  "point"         point NOT NULL,
  "polygon"       polygon NOT NULL,
  "real"          float4 NOT NULL,
  "smallint"      int2 NOT NULL,
  "smallserial"   serial2 NOT NULL,
  "serial"        serial4 NOT NULL,
  "text"          text NOT NULL,
  "time"          time [?][?] NOT NULL,
  "timestamp"     timestamp[?] [?] NOT NULL,
  "tsquery"       tsquery NOT NULL,
  "tsvector"      tsvector NOT NULL,
  "txid_snapshot" txid_snapshot NOT NULL,
  "uuid"          uuid NOT NULL,
  "xml"           xml NOT NULL
);