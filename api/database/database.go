package database

import (
	"database/sql"
	"fmt"
	"os"
)

// Init Initialize DB parameters and connection
func Init() (*sql.DB, error) {
	connstring := fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;database=%s",
		os.Getenv("SQL_SERVER"), os.Getenv("SQL_PORT"), os.Getenv("SQL_USER"), os.Getenv("SQL_PASS"), os.Getenv("SQL_DATABASE_NAME"))

	db, err := sql.Open("mssql", connstring)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return nil, err
	}

	return db, nil
}
