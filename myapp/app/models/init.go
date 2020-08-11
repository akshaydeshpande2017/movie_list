package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	host = "localhost"
	port = 5432
	user = "akshay"
	//password = "your-password"
	dbname = "akshay"
)

func InitDB() {
	connstring := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)

	var err error
	DB, err = sql.Open("postgres", connstring)
	if err != nil {
		fmt.Println("DB ERROR")
	}
	fmt.Println("DB Connected")
	_, err = DB.Query("CREATE TABLE IF NOT EXISTS movies (id serial PRIMARY KEY, name VARCHAR (50) UNIQUE NOT NULL, comment VARCHAR (50), rating INTEGER)")
	if err != nil {
		fmt.Println("Error Creating database table")
		DB.Close()
	}
}
