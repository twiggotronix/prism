package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	config "prism/proxy/config"
	db "prism/proxy/db"
	repository "prism/proxy/repository"

	controllers "prism/proxy/controllers"
)

func main() {
	envfile := flag.String("e", ".env", "path to the .env file")
	flag.Parse()

	redisClient := initRedis()
	conf := config.NewAppConfig(redisClient)
	conf.Init(envfile)

	dbConfig := config.NewDbConfig(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
	dbInstance := db.NewDb(dbConfig)
	defer dbInstance.Close()

	router := gin.Default()
	if os.Getenv("ENVIRONMENT") == "dev" {
		router.ForwardedByClientIP = true
		router.SetTrustedProxies([]string{"127.0.0.1"})
		router.Use(cors.Default())
	}

	proxyRepository := repository.NewProxyRepository(dbInstance.Get())
	appNetworkSettings := &config.AppNetworkSettings{
		Host: os.Getenv("HOST"),
		Ip:   os.Getenv("IP"),
	}
	proxyController := controllers.NewProxyController(proxyRepository, conf, *appNetworkSettings)
	router.GET("/api/proxies", proxyController.GetAll)
	router.GET("/api/proxies/:id", proxyController.Get)
	router.POST("/api/proxies", proxyController.Save)
	router.DELETE("/api/proxies/:id", proxyController.Delete)
	proxyController.InitProxies(router)

	configController := controllers.NewConfigController(conf)
	router.GET("/api/config", configController.Get)
	router.POST("/api/config", configController.Set)

	router.Run(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}

func initRedis() *redis.Client {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URI"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	ping, err := redisClient.Ping().Result()
	if err != nil {
		fmt.Printf("redis ping error: %s\n", err.Error())
	}
	fmt.Printf("redis ping: %s\n", ping)

	return redisClient
}
