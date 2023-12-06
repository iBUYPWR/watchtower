package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func database(dbconfig DatabaseConfig) *sql.DB {
	conn := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", dbconfig.User, dbconfig.Password, dbconfig.Name)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
	return db
}

func initTables(conn *sql.DB) {
	_, err := conn.Exec("CREATE TABLE Programs (id serial PRIMARY KEY, name VARCHAR(255), platform VARCHAR(255), submission VARCHAR(255), bounty BOOLEAN NOT NULL);")
	if err != nil {
		log.Print(fmt.Sprintf("[+] Info: %v", err))
	}
	_, err = conn.Exec("CREATE TABLE targets ( id serial PRIMARY KEY, name VARCHAR(10000), category VARCHAR(255), scope BOOLEAN NOT NULL, program_id INTEGER);")
	if err != nil {
		log.Print(fmt.Sprintf("[+] Info: %v", err))
	}
	_, err = conn.Exec("alter table targets add constraint fk_target foreign key (program_id) REFERENCES programs (id);")
	if err != nil {
		log.Print(fmt.Sprintf("[+] Info: %v", err))
	}
}
