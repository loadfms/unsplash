package services

import (
	"os"

	"github.com/loadfms/unsplash/models"
	"gopkg.in/yaml.v2"
)

func LoadConfig(configPath string) (*models.Configs, error) {
	config := &models.Configs{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
