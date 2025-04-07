package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

func Read() (*Config, error) {
	conf := &Config{}

	configPath, err := getConfigFile()
	if err != nil {
		return nil, err
	} else if configPath == nil {
		return nil, errors.New("invalid configuration path")
	}

	file, err := os.ReadFile(*configPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func (config Config) SetUser(userName string) error {
	config.CurrentUserName = userName

	err2 := writeToFile(config)
	if err2 != nil {
		return err2
	}

	return nil
}

// region PRIVATE
func getConfigFile() (*string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configPath := fmt.Sprintf("%s/%s", homeDir, configFileName)
	return &configPath, nil
}

func writeToFile(config Config) error {
	configPath, err := getConfigFile()
	if err != nil {
		return err
	} else if configPath == nil {
		return errors.New("invalid configuration path")
	}

	marshal, err := json.Marshal(config)
	if err != nil {
		return err
	} else if marshal == nil {
		return errors.New("could not marshal configuration")
	}

	err = os.WriteFile(*configPath, marshal, 0644)
	if err != nil {
		return err
	}

	return nil
}

//endregion PRIVATE
