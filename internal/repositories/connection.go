package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"golang_server.dankbueno.com/internal/config"
)

var db *sql.DB

func Connect() {

	cfg := mysql.Config{
		AllowNativePasswords: true,
		User:                 os.Getenv(config.DBUser),
		Passwd:               os.Getenv(config.DBPass),
		Net:                  "tcp",
		Addr:                 os.Getenv(config.DBHost),
		DBName:               os.Getenv(config.DBName),
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
