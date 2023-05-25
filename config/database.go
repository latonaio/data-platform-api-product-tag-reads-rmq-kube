package config

import (
	"fmt"
	"os"
)

type Database struct {
	user     string
	password string
	dbName   string
	address  string
	port     string
}

func newDatabase() *Database {
	return &Database{
		user:     os.Getenv("MONGODB_USER"),
		password: os.Getenv("MONGODB_PASSWORD"),
		dbName:   os.Getenv("MONGODB_DB_NAME"),
		address:  os.Getenv("DATA_PLATFORM_MONGODB_KUBE"),
		port:     os.Getenv("MONGODB_PORT"),
	}
}
func (c Database) MongodbDsn() string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		c.user, c.password, c.address, c.port,
	)
}

func (c Database) MongodbDsnDBName() string {
	return fmt.Sprintf("%s", c.dbName)
}
