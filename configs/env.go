package configs

import (
	"log"

	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API apiConfigs
	DB  dbConfigs
}

type apiConfigs struct {
	PORT string
}

type dbConfigs struct {
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

	cfg = new(config)

	cfg.API = apiConfigs{
		PORT: viper.GetString("PORT"),
	}
	cfg.DB = dbConfigs{
		URL: viper.GetString("DB_URL"),
	}
}

func GetDbURL() string {
	return cfg.DB.URL
}

func GetPort() string {
	return cfg.API.PORT
}
