#!/bin/bash
cd /
psql postgres -c "DROP DATABASE IF EXISTS postgres_db;"
psql postgres -c "DROP DATABASE IF EXISTS postgres_test_db;"
psql postgres -c "CREATE DATABASE postgres_db;"
psql postgres -c "CREATE DATABASE postgres_test_db;"
psql postgres -c "CREATE EXTENSION \"uuid-ossp\";" -d postgres_db
psql postgres -c "CREATE EXTENSION \"uuid-ossp\";" -d postgres_test_db
cd -
