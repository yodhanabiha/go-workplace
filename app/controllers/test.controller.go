package controllers

import (
	"encoding/json"
	"log"
	"nabiha/project-golang/app/config"
	"net/http"
)

type TestController struct {
	Config *config.Config
}

func NewTestController(cfg *config.Config) *TestController {
	return &TestController{Config: cfg}
}

func (uc *TestController) Route(w http.ResponseWriter, r *http.Request) {
	test, err := json.Marshal("Welcome to project Golang Nabiha!!")
	if err != nil {
		log.Printf("Error querying users: %v", err)
		return
	}
	w.Write(test)
}