package main

import (
	"fmt"
	"time"

	"github.com/rehabaam/TDBackend/config"
	labels "github.com/rehabaam/TDBackend/localization"
	"github.com/rehabaam/TDBackend/logger"
	commands "github.com/rehabaam/TDBackend/services"
)

// main func loads application's configurations and starts the HTTP server
func main() {

	// Get the time in RFC3339Milli format
	strTime := time.Now().Format(labels.RFC3339Milli)

	// Run HTTP server
	config.Load()

	// Init logger
	if err := logger.Init(config.AppConfig.Logging.LoggerLevel, strTime, config.AppConfig.Microservice.Processor); err != nil {
		fmt.Printf("failed to initialize logger: %v", err.Error())
	}
	_ = startServer()

}

// startServer func starts an HTTP server that exposes RESTful APIs
func startServer() error {
	logger.AppLogger(labels.Info, labels.ServerStarting, 0)
	err := commands.NewServer().Start()
	if err != nil {
		// Print the error occurred while starting the server.
		fmt.Printf("%+v\n", err.Error())
	}
	return err
}
