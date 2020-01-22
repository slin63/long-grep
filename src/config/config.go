package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type configParam struct {
	Address []string
}

// Addresses grabs the list of known addresses from our config.json.
func Addresses() ([]string, error) {
	configParams, err := parseJSON(os.Getenv("CONFIG"))
	if err != nil {
		return make([]string, 0), err
	}
	return configParams.Address, nil
}

func parseJSON(fileName string) (configParam, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return configParam{}, err
	}

	// Necessities for go to be able to read JSON
	fileString := string(file)

	fileReader := strings.NewReader(fileString)

	decoder := json.NewDecoder(fileReader)

	var configParams configParam

	// Finally decode into json object
	err = decoder.Decode(&configParams)
	if err != nil {
		return configParam{}, err
	}

	return configParams, nil
}
