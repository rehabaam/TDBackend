package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

var (
	configPath string = "config/prod.yml"
	// Config main config
	Config ActiveConfig
)

// ActiveConfig Micro-service's configs
type ActiveConfig struct {
	Microservice struct {
		Processor string `yaml:"processor" envconfig:"PROCESSOR"`
		Service   string `yaml:"service" envconfig:"SERVICE_NAME"`
	} `yaml:"microService"`
	Logging struct {
		LoggerLevel string `yaml:"level" envconfig:"LOGGER_LEVEL"`
	} `yaml:"logging"`
}

// Load config from config files or from environment variables
func Load() {

	// Read from yml file & Environment variables
	readFile(&Config, configPath)
	readEnv(&Config)

}

// Read yml file
func readFile(cfg interface{}, path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		panic(err)
	}
}

// Read from Environment variables
func readEnv(cfg interface{}) {
	err := envconfig.Process("", cfg)
	if err != nil {
		panic(err)
	}
}
