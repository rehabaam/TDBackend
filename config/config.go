package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

var (
	configPath string = "config/prod.yml"
	AppConfig  Config
)

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
		fmt.Println(err.Error())
	}
	// defer the closing of the file so that we can parse it later on
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(f)
	fmt.Println(decoder)
	if errD := decoder.Decode(cfg); errD != nil {
		fmt.Println(errD.Error())
	}
}

// Read from Environment variables
func readEnv(cfg interface{}) {
	err := envconfig.Process("", cfg)
	if err != nil {
		panic(err)
	}
}
