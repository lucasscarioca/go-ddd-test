package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/lucasscarioca/music-stash-server/configs"
)

var conn *sql.DB

func Connect() {
	var err error
	conn, err = sql.Open("postgres", configs.GetDbURL())
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := conn.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	// conn.SetConnMaxLifetime(time.Duration(10) * time.Second)
	// conn.SetMaxIdleConns(5)
	// conn.SetMaxOpenConns(2)

	fmt.Println("Connected to database...")
}

func GetDbConnection() *sql.DB {
	return conn
}
