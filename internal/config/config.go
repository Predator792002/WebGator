package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".gatorconfig.json"), nil
}

func Read() (Config, error) {
	path, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	jsonFile, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}

	defer jsonFile.Close()
	data, err := io.ReadAll(jsonFile)
	if err != nil {
		return Config{}, err
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	path, err := getConfigPath()
	if err != nil {
		return err
	}

	perm := os.FileMode(0600)
	if err = os.WriteFile(path, data, perm); err != nil {
		return err
	}
	return nil
}
