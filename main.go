package main

import (
	"github.com/Fachrulmustofa20/bank-example.git/config"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// @title           Document API Bank
// @version         1.0
// @description     Document API Bank
// @termsOfService  http://fachrulmustofa.site/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:45001
// @BasePath  /api/

// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.Init()

	err = cfg.Start()
	if err != nil {
		log.Fatal(err)
	}
}
