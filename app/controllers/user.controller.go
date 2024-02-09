package controllers

import (
	"encoding/json"
	"log"
	"nabiha/project-golang/app/config"
	"nabiha/project-golang/app/models"
	"net/http"
)

type UserController struct {
	Config *config.Config
}

func NewUserController(cfg *config.Config) *UserController {
	return &UserController{Config: cfg}
}

func (uc *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := uc.Config.DB.Pg.Query("SELECT * FROM users")
	if (err != nil) {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Age) // Adjust fields according to your user model
		if err != nil {
			log.Printf("Error querying users: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		log.Printf("Error querying users: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Encode users slice to JSON
	usersJSON, err := json.Marshal(users)
	if err != nil {
		log.Printf("Error querying users: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set response content type header
	w.Header().Set("Content-Type", "application/json")

	// Write response body with users JSON data
	w.Write(usersJSON)
}