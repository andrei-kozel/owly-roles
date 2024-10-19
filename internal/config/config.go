package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	ServiceName string `yaml:"service_name" required:"true"`
	Env         string `yaml:"env" required:"true"`
	PostgresUrl string `yaml:"postgres_url" required:"true"`
}

var AppConfig Config

func Configurations() (environment string, err error) {
	env := os.Getenv("ENV")

	if env == "" {
		env = "dev" // default to "dev" environment
	}

	viper.SetConfigName(env)         // Config file name without extension
	viper.SetConfigType("yaml")      // Config file format
	viper.AddConfigPath("./configs") // Path to look for the config file

	bindEnvironmentVariables()

	viper.SetEnvPrefix(strings.ToUpper(env)) // Optional: Set prefix based on environment

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("Config file not found or invalid: %v. Falling back to environment variables.", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
		return "", err
	}

	return env, nil
}

func bindEnvironmentVariables() {
	viper.BindEnv("PostgresUrl", "POSTGRES_URL")
	viper.BindEnv("ServiceName", "SERVICE_NAME")
	viper.BindEnv("Env", "ENV")
}
