package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	ServiceName     string `yaml:"service_name" required:"true"`
	Env             string `yaml:"env" required:"true"`
	PostgresUrl     string `yaml:"postgres_url" required:"true"`
	ApplicationPort int    `yaml:"application_port" required:"true"`
}

var AppConfig Config

func Configurations(path string) (environment string, err error) {
	env := os.Getenv("ENV")

	if env == "" {
		env = "dev" // default to "dev" environment
	}

	if path == "" {
		log.Printf("no config path provided. Using default path: ./configs")
		path = "./configs"
	}

	viper.SetConfigName(env)    // Config file name without extension
	viper.SetConfigType("yaml") // Config file format
	viper.AddConfigPath(path)   // Path to look for the config file

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
	viper.BindEnv("ApplicationPort", "APPLICATION_PORT")
}
