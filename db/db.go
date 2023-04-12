package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/lucasscarioca/music-stash-server/configs"
)

func OpenConnection() (*sql.DB, error) {
	dbConfigs := configs.GetDBEnv()

	sc := fmt.Sprintf("host=%s port =%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfigs.HOST, dbConfigs.PORT, dbConfigs.USER, dbConfigs.PASS, dbConfigs.DATABASE)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()

	// conn.SetConnMaxLifetime(time.Duration(10) * time.Second)
	// conn.SetMaxIdleConns(5)
	// conn.SetMaxOpenConns(2)

	return conn, err
}
