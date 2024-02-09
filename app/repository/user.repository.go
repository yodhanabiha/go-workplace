package repository

import (
	"log"
	"nabiha/project-golang/app/config"
	"nabiha/project-golang/app/models"
)

type UserRepository struct {
	DB *config.Config
}

func NewUserRepository(db *config.Config) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) FindAll() ([]models.User, error) {
	rows, err := ur.DB.DB.Pg.Query("SELECT * FROM users")
	if err != nil {
		log.Printf("Error querying users: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Age)
		if err != nil {
			log.Printf("Error scanning user row: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		log.Printf("Error iterating over user rows: %v", err)
		return nil, err
	}

	return users, nil
}
