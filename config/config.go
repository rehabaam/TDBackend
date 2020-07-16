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
	Microservice struct {
		Processor     string `yaml:"processor" envconfig:"MS_PROCESSOR"`
		MSServiceName string `yaml:"msServiceName" envconfig:"MS_SERVICE_NAME"`
	} `yaml:"microservice"`
	SSO struct {
		URL       string `yaml:"url" envconfig:"SSO_URL"`
		Username  string `yaml:"username" envconfig:"SSO_USERNAME"`
		Password  string `envconfig:"SSO_PASSWORD"`
		ClientID  string `envconfig:"SSO_CLIENT_ID"`
		Secret    string `envconfig:"SSO_SECRET"`
		Scope     string `yaml:"scope" envconfig:"SSO_SCOPE"`
		GrantType string `yaml:"grantType" envconfig:"SSO_GRANT_TYPE"`
	} `yaml:"sso"`
	Logging struct {
		LoggerLevel string `yaml:"loggerLevel" envconfig:"LOGGER_LEVEL"`
	} `yaml:"logging"`
	Database struct {
		HzlADDR string `yaml:"hzlADDR" envconfig:"HAZELCAST_ADDRESS"`
	} `yaml:"database"`
}

// Load config from config files or from environment variables
func Load() {

	// Get the time in RFC3339 format
	strTime := time.Now().Format(time.RFC3339)

	// Read from yml file & Environment variables
	readFile(&Serviceconfig, configPath)

	if err := logger.Init(Serviceconfig.Logging.LoggerLevel, strTime, Serviceconfig.Microservice.MSServiceName); err != nil {
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
