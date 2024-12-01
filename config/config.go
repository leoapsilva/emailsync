package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ConfigEnvVariable(key string) string {

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion: Key [%s] not found on environment variables (%s)", key, viper.GetViper().ConfigFileUsed())
	}

	return value
}
