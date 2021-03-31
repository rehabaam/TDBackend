package config

import (
	labels "TDBackend/localization"
	"TDBackend/logger"
	"fmt"
	"os"
	"time"

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

	// Get the time in RFC3339Milli format
	strTime := time.Now().Format(labels.RFC3339Milli)

	// Read from yml file & Environment variables
	readFile(&Config, configPath)

	if err := logger.Init(Config.Logging.LoggerLevel, strTime, Config.Microservice.Processor); err != nil {
		fmt.Printf("failed to initialize logger: %v", err.Error())
	}

	logger.Log.Debug(fmt.Sprintf("Default config file: '%s' is loaded...", configPath))
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
