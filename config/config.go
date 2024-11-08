package config

import (
	"os"

	"errors"

	"github.com/joho/godotenv"
)

func ReadEnviroment(fileName string) (connectionString, dirName string, err error) {
	err = godotenv.Load(fileName)
	if err != nil {
		return
	}

	connectionString = os.Getenv("CONNECTION_STRING")
	dirName = os.Getenv("DIR_NAME")
	if connectionString == "" || dirName == "" {
		return "", "", errors.New("CONNECTION_STRING or DIR_NAME not found")
	}

	return
}
