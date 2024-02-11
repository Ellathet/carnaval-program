package utils

import (
	"fmt"

	config "github.com/Ellathet/carnaval/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	env = config.Env
	DB  = initDatabase()
)

func initDatabase() *gorm.DB {

	db, err := gorm.Open(postgres.Open(getDSN()))
	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}

func getDSN() string {
	host := env.DatabaseHost
	port := env.DatabasePort
	user := env.DatabaseUsername
	password := env.DatabasePassword
	dbName := env.DatabaseName

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbName)
	return dsn
}
