package util

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// config for storing all configurations of the app
// values read by viper from a config file or env. variables
type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`      // Looks for app.env or app.yaml
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"` // Explicitly set to env format
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

// Loadconfig to read configuration from files or env variables
func LaodConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("failed to read config: %w", err)
	}

	err = viper.Unmarshal(&config)
	return config, err
}
