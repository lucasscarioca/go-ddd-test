package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/lucasscarioca/music-stash-server/configs"
)

var Conn *sql.DB

func Connect() {
	var err error
	dbConfig := configs.GetDBEnv()
	Conn, err = sql.Open("postgres", dbConfig.URL)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := Conn.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	// Conn.SetConnMaxLifetime(time.Duration(10) * time.Second)
	// Conn.SetMaxIdleConns(5)
	// Conn.SetMaxOpenConns(2)

	fmt.Println("Connected to database...")
}
