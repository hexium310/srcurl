package config

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

var ConfigFile string

func GetConfigFile() (string, error) {
	if ConfigFile == "" {
		return DefaultConfigFile()
	} else {
		return ConfigFile, nil
	}
}

func DefaultConfigFile() (string, error) {
	base, err := DefaultConfigDir()
	if err != nil {
		return "", err
	}
	filename := "srcurl.toml"

	configFile := path.Join(base, filename)
	return configFile, nil
}

func DefaultConfigDir() (string, error) {
	base, err := defaultConfigBaseDir()
	if err != nil {
		return "", err
	}
	rootDir := "srcurl"

	configDir := path.Join(base, rootDir)
	return configDir, nil
}

func defaultConfigBaseDir() (string, error) {
	dir, ok := os.LookupEnv("XDG_CONFIG_HOME")
	if ok {
		return dir, nil
	}

	switch runtime.GOOS {
	case "windows":
		env := "LOCALAPPDATA"
		dir, ok := os.LookupEnv(env)
		if !ok {
			return "", fmt.Errorf("%%%s%% not found", env)
		}

		return dir, nil
	default:
		env := "HOME"
		dir, ok := os.LookupEnv(env)
		if !ok {
			return "", fmt.Errorf("$%s not found", env)
		}
		dir = path.Join(dir, ".config")

		return dir, nil
	}
}
