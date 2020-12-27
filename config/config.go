package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// Configuration struct contains static info required to run the app
// It contains DB info
type Configuration struct {
	Address               string `env:"ADDRESS" envDefault:":8080"`
	JwtSecret             string `env:"JWT_SECRET,required"`
	DatabaseConnectionURL string `env:"CONNECTION_URL,required"`
	DatabaseName          string `env:"DATABASE_NAME,required"`
}

// Newconfig will read the data given from the .env file
func Newconfig(files ...string) *Configuration {
	err := godotenv.Load(files...) // Loading config file from env

	if err != nil {
		log.Print("No .env file was found %q\n", files)
	}

	cfg := Configuration{}

	// Parse env to configuration
	err = env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	return &cfg
}
