package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Attributes: []string{
			"--jinja", "--repeat-penalty", "1.0", "--sleep-idle-seconds", "45", "--host", "0.0.0.0",
		},
		Models: filepath.Join(userHome(), "Library", "Caches", "llama.cpp"),
	}

	cfgPath := filepath.Join(userHome(), ".config", "mm", "config.json")
	bs, err := os.ReadFile(cfgPath)

	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}

		return nil, err
	}

	err = json.Unmarshal(bs, cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func LoadRuntime() (*Runtime, error) {
	cfg := &Runtime{}

	cfgPath := filepath.Join(userHome(), ".config", "mm", "running.json")
	bs, err := os.ReadFile(cfgPath)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bs, cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func StoreRuntime(cfg *Runtime) error {
	cfgPath := filepath.Join(userHome(), ".config", "mm", "running.json")

	err := os.MkdirAll(filepath.Dir(cfgPath), 0755)

	if err != nil {
		return err
	}

	bs, err := json.MarshalIndent(cfg, "", "\t")

	if err != nil {
		return err
	}

	return os.WriteFile(cfgPath, bs, 0644)
}

func RemoveRuntime() error {
	cfgPath := filepath.Join(userHome(), ".config", "mm", "running.json")

	return os.Remove(cfgPath)
}

func userHome() string {
	home, err := os.UserHomeDir()

	if err != nil {
		return "."
	}

	return home
}
