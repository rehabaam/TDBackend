package main

import (
	"TDBackend/config"
	commands "TDBackend/services"
	"fmt"
)

func main() {

	// Run HTTP server
	config.Load()
	err := commands.RunServer()
	if err != nil {
		// Print the error occured while starting the server.
		fmt.Printf("%+v\n", err.Error())
	}
	return

}
