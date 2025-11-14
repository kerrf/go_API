package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "go_user:password@tcp(127.0.0.1:3306)/go_events?parseTime=true")

	if err != nil {
		// log.Fatal(err)
		panic("\x1b[31mCould not connect to database.\x1b[0m")
	}

	if err := DB.Ping(); err != nil {
		panic("Cannot connect to DB.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable :=
		`CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL
		)`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Something went wrong creating the users table...")
	}
	fmt.Println("Table 'users' is ready!")

	createEventsTable =
		`CREATE TABLE IF NOT EXISTS events (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description VARCHAR(255) NOT NULL,
		location VARCHAR(255) NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INT,
		FOREIGN KEY(user_id) REFERENCES users(id)
		)`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Something went wrong creating the events table")
	}
	fmt.Println("Table 'events' is ready!")

	createEventsTable =
		`CREATE TABLE IF NOT EXISTS registrations (
		id INT AUTO_INCREMENT PRIMARY KEY,
		event_id INT,
		user_id INT,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
		)`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Something went wrong creating the registrations table")
	}
	fmt.Println("Table 'registrations' is ready!")

}
