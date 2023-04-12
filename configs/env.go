package configs

import "github.com/spf13/viper"

var cfg *Config

type Config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	PORT string
}

type DBConfig struct {
	HOST     string
	PORT     string
	USER     string
	PASS     string
	DATABASE string
}

func setDefaults() {
	// SERVER
	viper.SetDefault("PORT", "3000")

	// DB
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_HOST", "localhost")
}

func Load() error {
	setDefaults()
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	cfg = new(Config)

	cfg.API = APIConfig{
		PORT: viper.GetString("PORT"),
	}
	cfg.DB = DBConfig{
		HOST:     viper.GetString("DB_HOST"),
		PORT:     viper.GetString("DB_PORT"),
		USER:     viper.GetString("DB_USER"),
		PASS:     viper.GetString("DB_PASS"),
		DATABASE: viper.GetString("DB_NAME"),
	}

	return nil
}

func GetDBEnv() DBConfig {
	return cfg.DB
}

func GetServerEnv() APIConfig {
	return cfg.API
}
