package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/nzmprlr/highway"
)

var (
	module, id, rev, builtAt string // injected at build time

	config = &Config{}
)

type Config struct {
	MaxLenFoo int `json:"maxLenFoo"`

	Highway *highway.Config `json:"highway"`
}

func Init() {
	file, err := os.Open(fmt.Sprintf(".config/%s.json", highway.Env()))
	if err != nil {
		panic(err)
	}

	defer file.Close()

	read, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(read, config)
	if err != nil {
		panic(err)
	}

	highway.App().Inject(module, id, rev, builtAt)
	highway.Bootstrap(config.Highway)
}

func Get() *Config {
	return config
}
