package core

import (
"github.com/caarlos0/env"
"github.com/joho/godotenv"
log "github.com/sirupsen/logrus"
)

type Settings struct {
	Bucket            string `env:"SHIELDBUCKET"`
}

var settings *Settings

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.WithFields(log.Fields{
			"envFile": true,
		}).Warn("Error loading .env file, reading configuration from ENV")
	}

	settings = &Settings{}
	err = env.Parse(settings)
	if err != nil {
		log.WithFields(log.Fields{
			"envFile": false,
		}).Fatal("Failed to parse ENV")
	}
}

func GetConfig() *Settings {
	return settings
}

