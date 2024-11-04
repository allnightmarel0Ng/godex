package config

import (
	"encoding/json"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GatewayPort      string
	ParserPort       string
	ContainerPort    string
	KafkaBroker      string
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	PostgresPort     string
	WhiteList        map[string]bool
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	whiteListStr := os.Getenv("WHITE_LIST")
	var whiteListSlice []string
	whiteList := make(map[string]bool)
	err = json.Unmarshal([]byte(whiteListStr), &whiteListSlice)
	if err != nil {
		return nil, err
	}

	for _, link := range whiteListSlice {
		whiteList[link] = true
	}

	return &Config{
		GatewayPort:      os.Getenv("GATEWAY_PORT"),
		ParserPort:       os.Getenv("PARSER_PORT"),
		ContainerPort:    os.Getenv("CONTAINER_PORT"),
		KafkaBroker:      os.Getenv("KAFKA_BROKER"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		WhiteList:        whiteList,
	}, nil
}
