#!/bin/bash
cd /
sudo -u postgres psql -c "DROP DATABASE IF EXISTS postgres_db;"
sudo -u postgres psql -c"DROP DATABASE IF EXISTS postgres_test_db;"
sudo -u postgres psql -c "CREATE DATABASE postgres_db;"
sudo -u postgres psql -c "CREATE DATABASE postgres_test_db;"
sudo -u postgres psql -c "CREATE EXTENSION \"uuid-ossp\";" -d postgres_db
sudo -u postgres psql -c "CREATE EXTENSION \"uuid-ossp\";" -d postgres_test_db
cd -
