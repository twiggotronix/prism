package config

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"

	models "prism/proxy/models"
)

type AppConfig interface {
	Init(envfile *string) error
	Get() models.Config
	Set(config models.Config) error
}
type ApplicationConfig struct {
	Client redis.Cmdable
}

type AppNetworkSettings struct {
	Host string
	Ip   string
}

func NewAppConfig(Client redis.Cmdable) AppConfig {
	return &ApplicationConfig{Client}
}

func (ac *ApplicationConfig) Init(envfile *string) error {
	log.Println("loading " + *envfile + " file ")
	err := godotenv.Load(*envfile)
	if err != nil {
		log.Fatal("Error loading " + *envfile + " file")
		return err
	}

	return nil
}

func (ac *ApplicationConfig) Get() models.Config {
	delay, err := ac.Client.Get("api-proxy-delay").Result()
	if err != nil {
		log.Printf("Failed to get delay from redis: %s\n", err.Error())
		log.Println("Defaulting to 0")
		delay = "0"
	}
	nbSecondsToWait, err := strconv.ParseInt(delay, 10, 64)
	if err != nil {
		log.Printf("Error converting delay config: %s\n", err)
		nbSecondsToWait = 0
	}
	prefix, err := ac.Client.Get("api-proxy-path-prefix").Result()
	if err != nil {
		log.Printf("Failed to get delay from redis: %s\n", err.Error())
		log.Println("Defaulting to /proxy/")
		prefix = "/proxy/"
	}

	return models.Config{
		Delay:       int(nbSecondsToWait),
		ProxyPrefix: prefix,
	}
}

func (ac *ApplicationConfig) Set(config models.Config) error {
	delayErr := ac.Client.Set("api-proxy-delay", config.Delay, time.Duration(0)).Err()
	if delayErr != nil {
		fmt.Printf("Error setting delay in redis: %s\n", delayErr)
		return delayErr
	}
	prefixErr := ac.Client.Set("api-proxy-path-prefix", config.ProxyPrefix, time.Duration(0)).Err()
	if prefixErr != nil {
		fmt.Printf("Error setting proxy prefix in redis: %s\n", prefixErr)
		return prefixErr
	}

	return nil
}

type DbConfig interface {
	GetUser() string
	GetPassword() string
	GetHost() string
	GetPort() string
	GetDatabaseName() string
}
type dbConfig struct {
	dbConfig models.DbConfig
}

func NewDbConfig(user string, password string, host string, port string, databaseName string) DbConfig {
	dbConfigModel := models.DbConfig{
		User:         user,
		Password:     password,
		Host:         host,
		Port:         port,
		DatabaseName: databaseName,
	}
	return &dbConfig{
		dbConfig: dbConfigModel,
	}
}

func (d *dbConfig) GetDatabaseName() string {
	return d.dbConfig.DatabaseName
}

func (d *dbConfig) GetHost() string {
	return d.dbConfig.Host
}

func (d *dbConfig) GetPassword() string {
	return d.dbConfig.Password
}

func (d *dbConfig) GetPort() string {
	return d.dbConfig.Port
}

func (d *dbConfig) GetUser() string {
	return d.dbConfig.User
}
