package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
)



var (
	db *gorm.DB
	once sync.Once
)

type Driver string

const (
	MySQL Driver = "mysql"
	Postgres Driver = "postgres"
)

type Config struct {
	Host string
	Port string
	User string
	Password string
	Database string
	Driver Driver
}

func DB(d Driver) {
	switch d {
	case MySQL:
		connectMysql(&Config{})
	case Postgres:
		connectPostgres(&Config{})
	}
}

func connectMysql(conf *Config) {
	once.Do(func() {
		var err error
		db, err = gorm.Open("mysql", conf.User+":"+conf.Password+"@tcp("+conf.Host+":"+conf.Port+")/"+conf.Database+"?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}
		fmt.Println("mysql connected")
	})
}

func connectPostgres(conf *Config) {
	once.Do(func() {
		var err error
		db, err = gorm.Open("postgres", conf.User+":"+conf.Password+"@tcp("+conf.Host+":"+conf.Port+")/"+conf.Database+"?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}
		fmt.Println("postgres connected")
	})
}

func GetDB() *gorm.DB {
	return db
}
