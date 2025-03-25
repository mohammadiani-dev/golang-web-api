package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Server struct {
	Port    string
	RunMode string
}

type Logger struct {
	FilePath string
	Encoding string
	Level    string
	Logger   string
}

type Redis struct {
	Host               string
	Port               string
	Password           string
	DB                 int
	PoolSize           int
	PoolTimeout        time.Duration
	IdleTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	IdleCheckFrequency time.Duration
}

type Postgres struct {
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}
type Config struct {
	Server   Server
	Redis    Redis
	Postgres Postgres
	Logger   Logger
}


func NewConfig() *Config {
	return &Config{}
}


func GetConfig() *Config {
	configPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := loadConfig(configPath, "yml")
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	config, err := parseConfig(v)
	if err != nil {
		log.Fatalf("error parsing config: %v", err)
	}

	return config
}

func parseConfig(v *viper.Viper) (*Config, error) {
	var config Config
	err := v.Unmarshal(&config)
	if err != nil {
		log.Printf("error unmarshalling config: %v", err)
		return nil, err
	}
	return &config, nil
}

func loadConfig(filename string, filetype string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.SetConfigType(filetype)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("error reading config file: %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, errors.New("error reading config file")
		}
		return nil, err
	}

	return v, nil
}

func getConfigPath(env string) string {
	if env == "production" {
		return "../config/config-production.yml"
	} else if env == "docker" {
		return "../config/config-docker.yml"
	} else {
		return "../config/config-development.yml"
	}

}
