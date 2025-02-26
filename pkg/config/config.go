package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	RabbitMQ RabbitMQConfig
}

type ServerConfig struct {
	Port string
}

type RabbitMQConfig struct {
	User      string
	Password  string
	Host      string
	Port      string
	URL       string
	Exchanges []string
}

var AppConfig Config

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	requiredVars := []string{"PORT", "RABBITMQ_USER", "RABBITMQ_PASSWORD", "RABBITMQ_HOST", "RABBITMQ_PORT"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Fatalf("Environment variable %s is required but not set", v)
		}
	}

	AppConfig = Config{
		Server: ServerConfig{
			Port: os.Getenv("PORT"),
		},
		RabbitMQ: RabbitMQConfig{
			User:      os.Getenv("RABBITMQ_USER"),
			Password:  os.Getenv("RABBITMQ_PASSWORD"),
			Host:      os.Getenv("RABBITMQ_HOST"),
			Port:      os.Getenv("RABBITMQ_PORT"),
			URL:       buildRabbitMQURL(),
			Exchanges: []string{"auth_exchange", "dlx_exchange"},
		},
	}
	log.Println("Configuration loaded successfully")
}

func buildRabbitMQURL() string {
	user := os.Getenv("RABBITMQ_USER")
	password := os.Getenv("RABBITMQ_PASSWORD")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")
	return "amqp://" + user + ":" + password + "@" + host + ":" + port
}
