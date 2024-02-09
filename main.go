package main

import (
	"log"
	"nabiha/project-golang/app/config"
	"nabiha/project-golang/app/controllers"
	"net/http"
)

func main() {
	// Load conf
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}

	// Init controllers
	userController := controllers.NewUserController(cfg)
	testController := controllers.NewTestController(cfg)

	// Define routes
	http.HandleFunc("/", testController.Route)
	http.HandleFunc("/users", userController.ListUsers)

	// Start server
	log.Printf("Server started on http://%s:%s",cfg.Server.Host, cfg.Server.Port)
	log.Fatal(http.ListenAndServe(cfg.Server.Host+":"+cfg.Server.Port, nil))
}