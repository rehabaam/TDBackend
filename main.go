package tdbackend

import (
	"TDBackend/config"
	commands "TDBackend/services"
	"fmt"
)

// main func loads application's configurations and starts the HTTP server
func main() {

	// Run HTTP server
	config.Load()
	startServer()

	return
}

// startServer func starts an HTTP server that exposes RESTful APIs
func startServer() error {
	err := commands.RunServer()
	if err != nil {
		// Print the error occured while starting the server.
		fmt.Printf("%+v\n", err.Error())
	}
	return err
}
