# Sample Go Web Service

This repository contains a web service written in Go that uses a PostgreSQL database.

It was designed to accommodate the following user story enumerating a list of requirements.

## Requirements

> * As a user I need an online address book exposed as a REST API.  I need the data set to include the following data fields: 
First Name, Last Name, Email Address, and Phone Number
> * I need the api to follow standard rest semantics to support listing entries, showing a specific single entry, and adding, modifying, and deleting entries.  
> * The code for the address book should include regular go test files that demonstrate how to exercise all operations of the service.  
> * Finally I need the service to provide endpoints that can export and import the address book data in a CSV format. 

## Database Design

To implement this design, I designed a simple database containing a single table that uses the following fields:
* id - a uuid used to uniquely identify address entries
* firstname
* lastname
* email
* phonenumber

I assumed that only id and email would be unique fields, as all the other fields could possibly be shared between people.

## Dependencies

Before running this project, there are a few dependencies that must be installed on your machine.

First of all, Go should be installed on your machine. I will assume this step has already been completed. If not, you can learn how to get started with go [here](https://golang.org/doc/install).

Go requires gcc, so it should be installed:
```shell
sudo apt-get install gcc
```

This project uses a makefile, so make should be installed. Here's how:
```shell
sudo apt-get install make
```

Next PostgreSQL must be installed. If you are using Ubuntu, it can be installed with the following script:
```shell
sudo apt-get install postgresql postgresql-contrib
```

Make sure PostgreSQL is running with this script:
```shell
sudo service postgresql start
```

Then configure the postgres user password to be 'postgres' which is what the project expects:
```shell
sudo -u postgres psql -c "ALTER USER postgres PASSWORD 'postgres';"
```

# Creating the Database

Next, use the _init-db.sh_ script at the root of the project to initialize the databases that this project uses, like so:
```shell
./init-db.sh
```
This can be run any time you feel like resetting this project's databases.

## Configuring the Environment

The environment must be configured before the server or any tests can be run.

Source the _env.sh_ script to configure GOPATH to include this project. This will also include env vars that will be used to access the database. It is a good idea to source env.sh before running any of the go code in this project.
```shell
source env.sh
```

Next, this project uses some go dependencies. The makefile has a list of all of the dependencies needed. To install them, run:
```shell
make deps
```

This project uses _goose_ to manage migrations. Run this to add migrations to the database:
```shell
make build
```
This will also build the project.

To run all tests in this project, run:
```shell
make test
```
The tests are set up to erase all db changes they have made after running.

To start a server instance, run:
```shell
make runserver
```
It will serve on localhost:8888. You can use Postman or cURL to call its APIs and make persistent changes to the database.

If you want to get rid of all your changes and start with a clean db, run _init_db.sh_ and you'll have a fresh database again. Keep in mind that you will have to run the migration again to re-initialize the address table.
