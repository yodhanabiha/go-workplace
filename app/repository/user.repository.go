package repository

import (
	"fmt"
	"log"
	"nabiha/project-golang/app/config"
	"nabiha/project-golang/app/models"
	"nabiha/project-golang/app/util"
	"reflect"
	"strconv"
)

type UserRepository struct {
	Config *config.Config
	Helper *util.HelperConfig
}

func NewUserRepository(config *config.Config) *UserRepository {
	return &UserRepository{Config: config, Helper: &util.HelperConfig{}}
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	query := "SELECT * FROM users"
	rows, err := r.Config.DB.Pg.Query(query)
	if err != nil {
		log.Printf("Error querying users: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Age, &user.At_created, &user.At_updated)
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

func (r *UserRepository) FindById(id string) (models.User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	var user models.User
	err := r.Config.DB.Pg.QueryRow(query, id).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Age, &user.At_created, &user.At_updated)
	if err != nil {
		log.Printf("Error querying users: %v", err)
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) FindOne(filter models.User) (models.User, error) {
	var user models.User
	var queryArgs []interface{}
	query := "SELECT * FROM users WHERE 1=1"

	filterValue := reflect.ValueOf(filter)
	for i := 0; i < filterValue.NumField(); i++ {
		fieldValue := filterValue.Field(i)
		if !reflect.ValueOf(fieldValue.Interface()).IsZero() {
			query += " AND " + filterValue.Type().Field(i).Name + " = $" + strconv.Itoa(len(queryArgs)+1)
			queryArgs = append(queryArgs, fieldValue.Interface())
		}
	}

	err := r.Config.DB.Pg.QueryRow(query, queryArgs...).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Age, &user.At_created, &user.At_updated)
	if err != nil {
		log.Printf("Error querying users: %v", err)
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Find(filter models.User) ([]models.User, error) {
	var users []models.User
	var queryArgs []interface{}
	query := "SELECT * FROM users WHERE 1=1"

	filterValue := reflect.ValueOf(filter)
	for i := 0; i < filterValue.NumField(); i++ {
		fieldValue := filterValue.Field(i)
		if !reflect.ValueOf(fieldValue.Interface()).IsZero() {
			query += " AND " + filterValue.Type().Field(i).Name + " = $" + strconv.Itoa(len(queryArgs)+1)
			queryArgs = append(queryArgs, fieldValue.Interface())
		}
	}

	rows, err := r.Config.DB.Pg.Query(query, queryArgs...)
	if err != nil {
		log.Printf("Error querying users: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Age, &user.At_created, &user.At_updated)
		if err != nil {
			log.Printf("Error scanning user: %v", err)
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Create(create models.User) (models.User, error) {
	query := "INSERT INTO users (username, email, password, age, at_created, at_updated) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, username, email, password, age"
	datenow := r.Helper.OnGetDateTimeNow()
	println(datenow)
	var user models.User
	err := r.Config.DB.Pg.
		QueryRow(query, create.Username, create.Email, create.Password, create.Age, datenow, datenow).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Age)
	if err != nil {
		log.Printf("Error: %v", err)
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) Update(update models.User, id string) (models.User, error) {
	query := "UPDATE users SET username=$1, age=$2, at_updated=$3 WHERE id=$4 RETURNING id, username, email, age"
	datenow := r.Helper.OnGetDateTimeNow()
	var user models.User
	err := r.Config.DB.Pg.
		QueryRow(query, update.Username, update.Age, datenow, id).
		Scan(&user.ID, &user.Username, &user.Email, &user.Age)
	if err != nil {
		log.Printf("Error: %v", err)
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) Delete(id string) (string, error) {
	query := "DELETE FROM users WHERE id=$1"
	_, err := r.Config.DB.Pg.Exec(query, id)
	if err != nil {
		log.Printf("Error: %v", err)
		return "", err
	}

	return fmt.Sprintf("succes delete data from id=%s", id), nil
}
