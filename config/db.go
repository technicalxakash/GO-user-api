package config

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {

	cfg := mysql.Config{
		User:      "root",
		Passwd:    "akash",
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    "usersdb",
		ParseTime: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Println("❌ ERROR: Failed to open database:", err)
		log.Fatal(err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Println("❌ ERROR: Failed to connect to database:", err)
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println(" Database connected successfully")

	return db
}
