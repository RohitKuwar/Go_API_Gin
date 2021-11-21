package config

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Port          string `mapstructure:"PORT"`
	FirestoreCred string `mapstructure:"FIRESTORE_CRED"`
	PrivateKey    string `mapstructure:"PRIVATE_KEY"`
	ProjectId     string `mapstructure:"PROJECT_ID"`
	ClientId      string `mapstructure:"CLIENT_ID"`
}
// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	// viper.SetConfigName("app")
	// viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
	
}
