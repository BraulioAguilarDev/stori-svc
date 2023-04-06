package config

import (
	"log"

	"github.com/spf13/viper"
)

type Settings struct {
	DSN_DB string
}

var Config Settings

func init() {
	viper.AutomaticEnv()
	viper.BindEnv("DSN_DB")

	viper.SetDefault("DSN_DB", "postgres://postgres:password@localhost:5432/stori_db?sslmode=disable")

	if err := viper.Unmarshal(&Config); err != nil {
		log.Panicf("Error unmarshalling configuration: %s", err)
	}

	log.Println("Parameters loaded are:")
	for _, k := range viper.AllKeys() {
		log.Printf("\t%s=%v\n", k, viper.Get(k))
	}
}
