package config

import (
	"encoding/json"
	"os"
)

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	script := &Config{}

	err = json.Unmarshal(data, script)
	if err != nil {
		return nil, err
	}

	return script, nil
}
