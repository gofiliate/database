package database

import (
	"database/sql" //sql lib
	"fmt"          //used for string manipulation here

	_ "github.com/go-sql-driver/mysql"
)

const connection_skeleton = "%s:%s@tcp(%s:%v)/%s" //format of mysql connection string

var ADPT *sql.DB //our database connection

// creates a formatted connection based on the supplied input
func GetConnectionString(user string, pass string, host string, port int, database string) (connectionString string) {

	connectionString = fmt.Sprintf(connection_skeleton, user, pass, host, port, database)
	return
}

// connects to the database defined by connection_string, returns an error if cannot connect
func Connect(connection_string string) (err error) {

	ADPT, err = sql.Open("mysql", connection_string)

	if err != nil {
		fmt.Printf("Error attempting to establish database connection: %s\n", err.Error())
	}

	err = ADPT.Ping()
	if err != nil {
		fmt.Printf("Error pinging database.ADPT: %v\n", err.Error())
	}
	return
}
