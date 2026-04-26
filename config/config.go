package config

import (
	"errors"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Logger   LoggerConfing
}

type ServerConfig struct {
	Port    string
	RunMode string
}
type LoggerConfing struct {
	FilePath   string
	Encoding   string
	Level      string
	LoggerName string
}
type PostgresConfig struct {
	Host                  string
	Port                  string
	User                  string
	Password              string
	DbName                string
	SSLMode               string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxLifetimeConnection time.Duration
}

type RedisConfig struct {
	Host               string
	Port               string
	User               string
	Password           string
	Db                 string
	SSLMode            string
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	PoolTimeout        time.Duration
	IdleCheckFrequency time.Duration
	PoolSize           int
}

func getConfigPath(appType string) string {
	switch appType {
	case "dev":
		return "dev_config.yaml"
	case "production":
		return "production_config.yaml"
	case "docker":
		return "docker_config.yaml"
	default:
		return ""
	}
}

func loadConfig(fileName, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(fileName)
	v.SetConfigType(fileType)
	v.AddConfigPath("./config")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	if err != nil {
		var configFileNotFoundErr *viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundErr) {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

func parseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func loadEnvs(filePath string) {
	err := godotenv.Load(filePath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetConfig() *Config {
	once.Do(func() {
		loadEnvs("./config/app.env")
		appType := os.Getenv("APP_TYPE")
		configType := os.Getenv("CONFIG_TYPE")
		if appType == "" || configType == "" {
			log.Fatal("APP_TYPE or CONFIG_TYPE environment variable not set")
		}
		cfgPath := getConfigPath(appType)
		viperCfg, err := loadConfig(cfgPath, configType)
		if err != nil {
			log.Fatal(err)
		}
		parsedCfg, err := parseConfig(viperCfg)
		if err != nil {
			log.Fatal(err)
		}
		config = parsedCfg
	})
	return config
}
