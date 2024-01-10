package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)



var (
	db *gorm.DB
	once sync.Once
)

type Driver string

const (
	MySQL Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

type Config struct {
	Host string
	Port string
	User string
	Password string
	Database string
	Driver Driver
}

func ConnectMysql(conf *Config) {
	once.Do(func() {
		var err error
		db, err = gorm.Open("mysql", conf.User+":"+conf.Password+"@tcp("+conf.Host+":"+conf.Port+")/"+conf.Database+"?parseTime=true")
		fmt.Println("mysql",conf.User+":"+conf.Password+"@tcp("+conf.Host+":"+conf.Port+")/"+conf.Database+"?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}
		fmt.Println("mysql connected")
	})
}

func ConnectPostgres(conf *Config) {
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
