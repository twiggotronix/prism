package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	config "prism/proxy/config"
	models "prism/proxy/models"
)

type Db interface {
	Close()
	Get() *gorm.DB
}

type dbInstance struct {
	Db *gorm.DB
}

func NewDb(dbConfig config.DbConfig) Db {
	db := initDb(dbConfig)
	return &dbInstance{
		Db: db,
	}
}

func initDb(dbConfig config.DbConfig) *gorm.DB {
	log.Println("initDb")
	dsn := dbConfig.GetUser() + ":" + dbConfig.GetPassword() + "@tcp(" + dbConfig.GetHost() + ":" + dbConfig.GetPort() + ")/" + dbConfig.GetDatabaseName() + "?charset=utf8mb4&parseTime=True&loc=Local"
	dbi, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("AutoMigrate")
	dbi.AutoMigrate(&models.Proxy{})

	return dbi
}

func (d *dbInstance) Get() *gorm.DB {
	return d.Db
}

func (d *dbInstance) Close() {
	log.Println("closing db connection")
	dbi, err := d.Db.DB()
	if err != nil {
		log.Println("Could not close db connection")
		return
	}
	dbi.Close()
}
