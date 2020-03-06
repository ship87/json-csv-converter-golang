package models

import (
	"github.com/joho/godotenv"
	"github.com/ship87/json-csv-converter-golang/helpers"
	"log"
	"os"
	"reflect"
)

var envVariables = map[string]string{
	"DirectoryDownload": "JSON_CSV_CONVERTER_GOLANG_DIRECTORY_DOWNLOAD",
	"PrefixFile":        "JSON_CSV_CONVERTER_GOLANG_PREFIX_FILE",
	"AppUrl":            "JSON_CSV_CONVERTER_GOLANG_APP_URL",
	"AppPort":           "JSON_CSV_CONVERTER_GOLANG_APP_PORT",
}

// Config - configuration of app
type Config struct {
	DirectoryDownload string
	PrefixFile        string
	AppUrl            string
	AppPort           string
}

func init() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// Fill - filling the configuration with current data
func (c *Config) Fill() {

	for index, variable := range envVariables {

		value, exists := os.LookupEnv(variable)
		if !exists {
			errorMessage := helpers.ConcatStrings([]string{"Not exist env variable ", variable})
			log.Fatal(errorMessage)
		}
		if value == "" {
			errorMessage := helpers.ConcatStrings([]string{"Env variable ", value, " not have value"})
			log.Fatal(errorMessage)
		}

		reflect.ValueOf(c).Elem().FieldByName(index).SetString(value)
	}
}
