package configs

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	KafkaBroker string
	ServerPort  string
)

func Load() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	kafkaBroker := os.Getenv("KAFKA_BROKER")
	if kafkaBroker == "" {
		kafkaBroker = "localhost:9092"
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8085"
	}

}
