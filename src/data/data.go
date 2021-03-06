package data

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbhost     = "DBHOST"
	dbport     = "DBPORT"
	dbuser     = "DBUSER"
	dbpass     = "DBPASS"
	dbname     = "DBNAME"
	testdbname = "TESTDBNAME"
)

// set up app db connection
func InitDb() func() {
	dbConfig := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig[dbhost], dbConfig[dbport], dbConfig[dbuser], dbConfig[dbpass], dbConfig[dbname])

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return func() {
		// defer close db
		db.Close()
	}
}

// set up test db connection
func InitTestDb() func() {
	dbConfig := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig[dbhost], dbConfig[dbport], dbConfig[dbuser], dbConfig[dbpass], dbConfig[dbname])

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return func() {
		// clean up test db
		statementStr := `TRUNCATE address CASCADE`
		_, err := db.Exec(statementStr)
		if err != nil {
			panic(err)
		}
	}
}

// load db config values from env
func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhost)
	checkVarOk(dbhost, ok)
	port, ok := os.LookupEnv(dbport)
	checkVarOk(dbport, ok)
	user, ok := os.LookupEnv(dbuser)
	checkVarOk(dbuser, ok)
	password, ok := os.LookupEnv(dbpass)
	checkVarOk(dbpass, ok)
	name, ok := os.LookupEnv(dbname)
	checkVarOk(dbname, ok)
	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name
	return conf
}

// load test db config values from env
func dbTestConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhost)
	checkVarOk(dbhost, ok)
	port, ok := os.LookupEnv(dbport)
	checkVarOk(dbport, ok)
	user, ok := os.LookupEnv(dbuser)
	checkVarOk(dbuser, ok)
	password, ok := os.LookupEnv(dbpass)
	checkVarOk(dbpass, ok)
	name, ok := os.LookupEnv(testdbname)
	checkVarOk(dbname, ok)
	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name
	return conf
}

func checkVarOk(varName string, ok bool) {
	if !ok {
		panic(fmt.Sprintf("%s environment variable required but not set", varName))
	}
}
