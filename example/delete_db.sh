#!/bin/bash

psql --command "DROP DATABASE IF EXISTS testing;"
psql --command "DROP USER IF EXISTS testing;"