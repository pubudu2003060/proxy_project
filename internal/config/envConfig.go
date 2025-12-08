package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Port  string
	DBURL string
}

func NewEnvConfig() EnvConfig {
	appEnv := os.Getenv("APP_ENV")
	log.Printf("APP ENV : %v", appEnv)

	if strings.TrimSpace(strings.ToLower(appEnv)) == "dev" || strings.TrimSpace(strings.ToLower(appEnv)) == "" {
		if err := godotenv.Load(".env.dev"); err != nil {
			log.Fatal("Error in loading .env.dev file:", err)
		}
	} else {
		log.Println("Running in production mode, using environment variables from Docker")
	}

	config := EnvConfig{
		Port:  getEnv("PORT", "8080"),
		DBURL: getEnv("DB_URL", ""),
	}

	config.validate()

	return config
}

func getEnv(name string, default_name string) string {
	if env, ok := os.LookupEnv(name); ok {
		return env
	}
	return default_name
}

func (e *EnvConfig) validate() {
	isErr := false
	message := ""

	if e.DBURL == "" {
		isErr = true
		message += "Empty database URL, "
	}

	if isErr {
		log.Fatal("error in env: ", message)
	}

}
