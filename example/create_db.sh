#!/bin/bash

psql --command "CREATE USER testing WITH SUPERUSER PASSWORD 'testing';"

createdb -O testing testing
psql -d testing -c "CREATE EXTENSION IF NOT EXISTS citext;"

psql testing -f ./schema.sql

go run generate.go