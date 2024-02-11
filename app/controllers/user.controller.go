package controllers

import (
	"log"
	"nabiha/project-golang/app/config"
	"nabiha/project-golang/app/models"
	"nabiha/project-golang/app/repository"
	"nabiha/project-golang/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserRepo *repository.UserRepository
}

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserController(cfg *config.Config) {
	controller := &UserController{
		UserRepo: repository.NewUserRepository(cfg),
	}

	cfg.Router.POST("/login", func(c *gin.Context) {
		var auth Auth
		err := c.ShouldBindJSON(&auth)
		if err != nil {
			log.Printf("Error decoding request body: %v", err)
			c.IndentedJSON(http.StatusBadRequest, "Bad request")
			return
		}

		filter := models.User{
			Email: auth.Email,
		}

		user, err := controller.UserRepo.FindOne(filter)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, "User not found!")
			return
		}

		checkPw := services.VerifyPassword(user.Password, auth.Password)
		if checkPw != nil {
			c.IndentedJSON(http.StatusNotAcceptable, "wrong password!")
			return
		}

		tokenString, err := services.CreateToken(user.Email)
		if err != nil {
			log.Printf("Error: %v", err)
			c.IndentedJSON(http.StatusInternalServerError, "Internal server error")
			return
		}
		c.IndentedJSON(http.StatusOK, tokenString)
	})

	cfg.Router.POST("/register", func(c *gin.Context) {
		var newUser models.User
		err := c.ShouldBindJSON(&newUser)
		if err != nil {
			log.Printf("Error decoding request body: %v", err)
			c.IndentedJSON(http.StatusBadRequest, "Bad request")
			return
		}

		hashedPassword, err := services.HashPassword(newUser.Password)
		if err != nil {
			log.Printf("Error hashing password: %v", err)
			c.IndentedJSON(http.StatusInternalServerError, "Internal server error")
			return
		}
		newUser.Password = hashedPassword

		created, err := controller.UserRepo.Create(newUser)
		if err != nil {
			log.Printf("Error: %v", err)
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
			return
		}
		c.IndentedJSON(http.StatusCreated, created)
	})

	cfg.Router.GET("/profile", func(c *gin.Context) {
		email, err := services.ProtectedHandler(c)
		if err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
			return
		}
		filter := models.User{
			Email: email,
		}
		user, err := controller.UserRepo.FindOne(filter)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, "User not found!")
			return
		}
		c.IndentedJSON(http.StatusOK, user)

	})

	cfg.Router.GET("/users", func(c *gin.Context) {
		_, err := services.ProtectedHandler(c)
		if err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
			return
		}
		users, err := controller.UserRepo.FindAll()
		if err != nil {
			log.Printf("Error listing users: %v", err)
			c.IndentedJSON(http.StatusInternalServerError, "Internal server error")
			return
		}
		c.IndentedJSON(http.StatusOK, users)
	})

	cfg.Router.GET("/user", func(c *gin.Context) {
		id := c.Query("id")
		if id == "" {
			c.IndentedJSON(http.StatusNotAcceptable, "no params connstrain id")
			return
		}

		user, err := controller.UserRepo.FindById(id)
		if err != nil {
			log.Printf("Error listing users: %v", err)
			c.IndentedJSON(http.StatusInternalServerError, "Internal server error")
			return
		}

		c.IndentedJSON(http.StatusOK, user)
	})

	cfg.Router.PUT("/users", func(c *gin.Context) {

		id := c.Query("id")
		if id == "" {
			c.IndentedJSON(http.StatusNotAcceptable, "no params connstrain id")
			return
		}
		var updateUser models.User
		err := c.ShouldBindJSON(&updateUser)
		if err != nil {
			log.Printf("Error decoding request body: %v", err)
			c.IndentedJSON(http.StatusBadRequest, "Bad request")
			return
		}

		updated, err := controller.UserRepo.Update(updateUser, id)
		if err != nil {
			log.Printf("Error: %v", err)
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
			return
		}

		c.IndentedJSON(http.StatusCreated, updated)

	})

	cfg.Router.DELETE("/users", func(c *gin.Context) {

		id := c.Query("id")
		if id == "" {
			c.IndentedJSON(http.StatusNotAcceptable, "no params connstrain id")
			return
		}

		deleted, err := controller.UserRepo.Delete(id)
		if err != nil {
			log.Printf("Error: %v", err)
			c.IndentedJSON(http.StatusInternalServerError, "Internal server error")
			return
		}

		c.IndentedJSON(http.StatusOK, deleted)

	})

}
