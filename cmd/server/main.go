package main

import (
	"log"

	"github.com/pubudu2003060/proxy_project/internal/config"
)

func main() {
	envConfig := config.NewEnvConfig()
	log.Println("server run in port:", envConfig.Port)
}
