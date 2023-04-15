package configs

import (
	"log"

	"github.com/spf13/viper"
)

var cfg *Config

type Config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	PORT string
}

type DBConfig struct {
	URL string
}

func setDefaults() {
	// SERVER
	viper.SetDefault("PORT", "3000")
	// DB
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_HOST", "localhost")
}

func Load() {
	setDefaults()
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to load environment variables: " + err.Error())
	}

	cfg = new(Config)

	cfg.API = APIConfig{
		PORT: viper.GetString("PORT"),
	}
	cfg.DB = DBConfig{
		URL: viper.GetString("DB_URL"),
	}
}

func GetDBEnv() DBConfig {
	return cfg.DB
}

func GetServerEnv() APIConfig {
	return cfg.API
}
