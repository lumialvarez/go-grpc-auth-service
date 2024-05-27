package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

//goland:noinspection SpellCheckingInspection
type Config struct {
	Port               string `mapstructure:"PORT"`
	JwtExpirationHours int64  `mapstructure:"JWT_EXPIRATION_HOURS"`
	JwtIssuer          string `mapstructure:"JWT_ISSUER"`
	DBUrl              string
	RabbitMQUrl        string
	JwtSecretKey       string
}

//goland:noinspection SpellCheckingInspection
func LoadConfig() (c Config, err error) {
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

	err = viper.Unmarshal(&c)

	c.DBUrl = readEnvironmentVariable("DB_URL")
	c.RabbitMQUrl = readEnvironmentVariable("RABBITMQ_URL")
	c.JwtSecretKey = readEnvironmentVariable("JWT_SECRET_KEY")

	return
}

func readEnvironmentVariable(envName string) string {
	envValue := os.Getenv(envName)
	if len(envValue) == 0 {
		panic(fmt.Sprintf("Environment Variable %s is not set", envName))
	}
	return envValue
}
