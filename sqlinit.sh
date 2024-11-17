#!/bin/bash
echo "Creating database role: ${DB_USER}"
set -e

POSTGRES="psql --username ${POSTGRES_USER} --dbname=${POSTGRES_DB}"

$POSTGRES <<-EOSQL
CREATE USER ${DB_RO_USER} WITH PASSWORD '${DB_RO_PASSWORD}';
GRANT pg_read_all_data TO ${DB_RO_USER};
ALTER SYSTEM SET wal_level = logical;
EOSQL
