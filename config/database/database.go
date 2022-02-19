package database

import "os"

func GetUri() string{
	return os.Getenv("DB_URL")
}

func GetDatabaseName() string{
	return os.Getenv("DB_NAME")
}