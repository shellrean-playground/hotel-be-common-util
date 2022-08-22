package configuration

import (
	"flag"
	"fmt"
	"github.com/shellrean-playground/hotel-be-common-util/config"
	"gopkg.in/yaml.v2"
	"os"
)

func New(configPath string) (*config.Config, error) {
	configImpl := &config.Config{}
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	d := yaml.NewDecoder(file)
	if err := d.Decode(&configImpl); err != nil {
		return nil, err
	}

	return configImpl, nil
}

func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}

	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}

	return nil
}

func ParseFlags() (string, error) {
	var configPath string

	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")
	flag.Parse()
	if err := validateConfigPath(configPath); err != nil {
		return "", err
	}
	return configPath, nil
}
