package config

import "github.com/spf13/viper"

//goland:noinspection SpellCheckingInspection
type Config struct {
	Port         string `mapstructure:"PORT"`
	DBUrl        string `mapstructure:"DB_URL"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

//goland:noinspection SpellCheckingInspection
func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./src/cmd/devapi/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
