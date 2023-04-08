package config

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	DSN_DB    string
	SG_KEY    string
	SG_SENDER string
	CSV_FILE  string
}

var Config config

func init() {
	viper.AutomaticEnv()
	viper.BindEnv("DSN_DB")
	viper.BindEnv("SG_KEY")
	viper.BindEnv("SG_SENDER")
	viper.BindEnv("CSV_FILE")

	viper.SetDefault("CSV_FILE", "data/txns.csv")

	if err := viper.Unmarshal(&Config); err != nil {
		log.Panicf("Error unmarshalling configuration: %s", err)
	}

	log.Println("Parameters loaded are:")
	for _, k := range viper.AllKeys() {
		log.Printf("\t%s=%v\n", k, viper.Get(k))
	}
}
