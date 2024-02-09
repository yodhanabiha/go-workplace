package controllers

import (
	"encoding/json"
	"log"
	"nabiha/project-golang/app/config"
	"nabiha/project-golang/app/repository"
	"net/http"
)

type UserController struct {
	UserRepo *repository.UserRepository
}

func NewUserController(db *config.Config) *UserController {
	return &UserController{
		UserRepo: repository.NewUserRepository(db),
	}
}

func (uc *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.UserRepo.FindAll()
	if err != nil {
		log.Printf("Error listing users: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	usersJSON, err := json.Marshal(users)
	if err != nil {
		log.Printf("Error marshalling users to JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(usersJSON)
}
