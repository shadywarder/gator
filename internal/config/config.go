package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

const (
	configFileName = ".gatorconfig.json"
)

// Config struct represents the structure of the .gatorconfig.
type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// MustLoad instantiates new Config entity.
func MustLoad() (*Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var cfg Config

	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// SetUser writes config struct to the JSON.
func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user
	return write(*c)
}

// getConfigFilePath constructs config path.
func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, configFileName), nil
}

// write writes updates config to the JSON.
func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(cfg); err != nil {
		return err
	}

	return nil
}
