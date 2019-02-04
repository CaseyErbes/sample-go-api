export PATH="$PATH:$(go env GOPATH)/bin"
export GOPATH="$(go env GOPATH):$(pwd)"
export GOBIN="$(pwd)/bin"

export DBHOST=localhost
export DBPORT=5432
export DBUSER=postgres
export DBPASS=postgres
export DBNAME=postgres_db
