package config

import (
	"encoding/json"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	EntrypointPort   string
	ParserPort       string
	ContainerPort    string
	KafkaBroker      string
	PostgresName     string
	PostgresUser     string
	PostgresPassword string
	PostgresPort     string
	WhiteList        []string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	whiteListStr := os.Getenv("WHITE_LIST")
	var whiteList []string
	err = json.Unmarshal([]byte(whiteListStr), &whiteList)
	if err != nil {
		return nil, err
	}

	return &Config{
		EntrypointPort:   os.Getenv("ENTRYPOINT_PORT"),
		ParserPort:       os.Getenv("PARSER_PORT"),
		ContainerPort:    os.Getenv("CONTAINER_PORT"),
		KafkaBroker:      os.Getenv("KAFKA_BROKER"),
		PostgresName:     os.Getenv("POSTGRES_NAME"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		WhiteList:        whiteList,
	}, nil
}
