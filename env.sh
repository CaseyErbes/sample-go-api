#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

export PATH="$PATH:$(go env GOPATH)/bin"
export GOPATH="$(go env GOPATH):$DIR"
export GOBIN="$DIR/bin"

export DBHOST=localhost
export DBPORT=5432
export DBUSER=postgres
export DBPASS=postgres
export DBNAME=postgres_db
export TESTDBNAME=postgres_test_db
