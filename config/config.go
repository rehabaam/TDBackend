package config

import (
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

var (
	configPath string = "config/prod.yml"
	// Config main config
	AppConfig Config
)

// ActiveConfig Micro-service's configs
type Config struct {
	Microservice struct {
		Processor string `yaml:"processor" envconfig:"PROCESSOR_NAME"`
		Service   string `yaml:"service" envconfig:"SERVICE_NAME"`
	} `yaml:"microservice"`
	Logging struct {
		LoggerLevel string `yaml:"level" envconfig:"LOGGER_LEVEL"`
	} `yaml:"logging"`
}

// Load config from config files or from environment variables
func Load() {

	// Read from yml file & Environment variables
	readFile(&AppConfig, configPath)
	readEnv(&AppConfig)

}

// Read yml file
func readFile(cfg interface{}, path string) {
	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		panic(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

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
