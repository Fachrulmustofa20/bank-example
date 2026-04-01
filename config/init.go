package config

import log "github.com/sirupsen/logrus"

func Init() Config {
	var cfg Config
	err := cfg.initPostgres()
	if err != nil {
		log.Panic()
	}

	initLogger()
	log.Info("Server is running ..")

	return cfg
}
