package main

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

const (
	DBHost     = "localhost"
	DBPort     = 5432
	DBUser     = "postgres"
	DBPassword = "postgres"
	DBname     = "postgres"
)

func initDB() {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBPassword, DBname)

	if DB, err = sql.Open("postgres", psqlInfo); err != nil {
		panic(fmt.Errorf("failed to open database: %s, %v", DBname, err))
	}

	if err = DB.Ping(); err != nil {
		panic(fmt.Errorf("failed to ping database: %s, %v", DBname, err))
	} else {
		fmt.Println("Successfully ping db")
	}
}
