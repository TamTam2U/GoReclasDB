package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	Host string
	Port string
	User string
	Password string
	Database string
	Driver Driver
}

func TestDb(t *testing.T) {
	config := Config{
		Host: "localhost",
		Port: "3306",
		User: "root",
		Password: "",
		Database: "butitin-test",
		Driver: MySQL,
	}
	
	ConnectMysql(&config)
	GetDB()
	assert.Equal(t, GetDB(), GetDB())
}

func BenchmarkDb(b *testing.B) {
	config := Config{
		Host: "localhost",
		Port: "3306",
		User: "root",
		Password: "",
		Database: "butitin-test",
		Driver: MySQL,
	}
	
	ConnectMysql(&config)
	GetDB()
	assert.Equal(b, GetDB(), GetDB())
}