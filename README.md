# DAO Pattern in Go

## Structure / Architecture

### cmd/
Main package. Holds the main function.

### dao/
This folder contains the factory, repository and folders for each DB.

### repository/
The repository holds interfaces for each model (e.g. UserDao interface for the User Model)
Those interfaces then define the methods needed (e.g. create, getAll, etc.)

### [dbname]/
Each DB has its own folder which contains a connection.go and implementation file for each interface.
The connection.go file opens the DB connection and returns the DB reference.
The implementation files implement the methods defined in the repository interfaces specifically for the DB.

### factory/
The factory returns the repository.DAO for whatever DB specified as a parameter.
