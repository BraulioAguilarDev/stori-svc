package config

import (
	"log"

	"github.com/jinzhu/configor"
)

type Sngrid struct {
	Key    string
	Sender string
	Name   string
}

type config struct {
	DSN_DB    string `default:"postgres://postgres:postgres@localhost:5432/dbname?sslmode=disable" env:"DSN_DB"`
	SG_KEY    string `default:"token" env:"SG_KEY"`
	SG_SENDER string `default:"info@stori.com" env:"SG_SENDER"`
}

var Config config

func init() {
	if err := configor.Load(&Config); err != nil {
		log.Fatal(err.Error())
	}
}
