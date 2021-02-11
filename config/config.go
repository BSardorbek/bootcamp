package config

import (
	"database/sql"
	"fmt"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
	db       *sql.DB
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "admin",
		Password: "xoji",
		DBName:   "mydb",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
	)
}
