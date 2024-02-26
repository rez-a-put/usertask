package utils

import (
	"log"
	"os"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func init() {
	projectDirName := "usertask"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		os.Exit(1)
	}
}

func GetEnvByKey(key string) string {
	return os.Getenv(key)
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Please input right format of email"
	case "max":
		return "Maximum " + fe.Param() + " characters"
	case "min":
		return "Minimum " + fe.Param() + " characters"
	}
	return "Unknown error"
}
