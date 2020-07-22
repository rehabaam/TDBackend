package config

import (
	"TDBackend/logger"
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

var (
	configPath string = "config/prod.yml"
	// Serviceconfig main config
	Serviceconfig ActiveConfig
)

// ActiveConfig Micro-service's configs
type ActiveConfig struct {
	Service struct {
		Name string `yaml:"name" envconfig:"NAME"`
	} `yaml:"service"`
	Logging struct {
		LoggerLevel string `yaml:"loggerLevel" envconfig:"LOGGER_LEVEL"`
	} `yaml:"logging"`
}

// Load config from config files or from environment variables
func Load() {

	// Get the time in RFC3339 format
	strTime := time.Now().Format(time.RFC3339)

	// Read from yml file & Environment variables
	readFile(&Serviceconfig, configPath)

	if err := logger.Init(Serviceconfig.Logging.LoggerLevel, strTime, Serviceconfig.Service.Name); err != nil {
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
