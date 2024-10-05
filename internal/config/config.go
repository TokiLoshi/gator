package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DBUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user"`
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return homeDir, fmt.Errorf("error getting home path file: %w", err)
	}
	path := filepath.Join(homeDir, configFileName)

	return path, nil
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error generating file path name: %w", err)
	}

	file, err := os.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("error opening file: %w", err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)

	if err != nil {
		return Config{}, fmt.Errorf("error reading file, %w", err)
	}

	return cfg, nil
}

func (cfg *Config) SetUser(user string) error {
	cfg.CurrentUserName = user 
	return write(*cfg)
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting file path")
	}
	
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
