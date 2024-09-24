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

const configFileName = "/workspace/.gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return homeDir, fmt.Errorf("error getting home path file: %w", err)
	}
	fmt.Printf("home directory: %v\n", homeDir)
	path := filepath.Join(homeDir, configFileName)

	return path, nil
}

func Read() (Config, error) {
	fmt.Println("Commence Reading...")

	var cfg Config
	path, err := getConfigFilePath()
	if err != nil {
		return cfg, fmt.Errorf("error generating file path name: %w", err)
	}

	fmt.Printf("resutling file: %v\n", path)
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Blunder here: %v\n", err)
		return cfg, fmt.Errorf("error opening file: %w", err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)

	if err != nil {
		return cfg, fmt.Errorf("error reading file, %w", err)
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
