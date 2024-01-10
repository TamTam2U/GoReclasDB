package database

import (
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

