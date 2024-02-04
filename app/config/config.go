package config

import (
	"bitohw_xin/app/server"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server *server.Config
}

func Get(path string) (*Config, error) {
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	config := new(Config)
	err = yaml.Unmarshal(bytes, config)

	if err != nil {
		return nil, err
	}

	return config, nil
}
