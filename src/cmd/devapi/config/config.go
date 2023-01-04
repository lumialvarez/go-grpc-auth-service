package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

//goland:noinspection SpellCheckingInspection
type Config struct {
	Port         string `mapstructure:"PORT"`
	DBUrl        string `mapstructure:"DB_URL"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

//goland:noinspection SpellCheckingInspection
func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./src/cmd/devapi/config/envs")

	scope := os.Getenv("SCOPE")
	log.Print("OS scope configuration:", scope)
	if len(scope) == 0 {
		scope = "dev"
	}
	log.Print("Selected configuration:", scope)
	viper.SetConfigName(scope)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
