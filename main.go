package main

import (
	"fmt"
	"log"
	"nabiha/project-golang/app/config"
	"nabiha/project-golang/app/controllers"
)

func main() {

	// Load conf
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}
	//setup gin
	host := cfg.Server.Host
	port := cfg.Server.Port
	address := fmt.Sprintf("%s:%s", host, port)

	// Init controllers
	controllers.NewUserController(cfg)
	controllers.NewTestController(cfg)

	// Start server
	cfg.Router.Run(address)
	log.Printf("Server started on http://%s:%s", host, port)
	// log.Fatal(http.ListenAndServe(cfg.Server.Host+":"+cfg.Server.Port, nil))
}
