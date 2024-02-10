package controllers

import (
	"log"
	"nabiha/project-golang/app/config"
	"nabiha/project-golang/app/models"
	"nabiha/project-golang/app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserRepo *repository.UserRepository
}

func NewUserController(cfg *config.Config) {
	controller := &UserController{
		UserRepo: repository.NewUserRepository(cfg),
	}

	cfg.Router.GET("/users", func(c *gin.Context) {
		users, err := controller.UserRepo.FindAll()
		if err != nil {
			log.Printf("Error listing users: %v", err)
			c.IndentedJSON(http.StatusInternalServerError, "Internal server error")
			return
		}
		c.IndentedJSON(http.StatusAccepted, users)
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

		c.IndentedJSON(http.StatusAccepted, user)
	})

	cfg.Router.POST("/users", func(c *gin.Context) {

		var newUser models.User
		err := c.ShouldBindJSON(&newUser)
		if err != nil {
			log.Printf("Error decoding request body: %v", err)
			c.IndentedJSON(http.StatusBadRequest, "Bad request")
			return
		}

		created, err := controller.UserRepo.Create(newUser)
		if err != nil {
			log.Printf("Error: %v", err)
			c.IndentedJSON(http.StatusInternalServerError, "Internal server error")
			return
		}
		c.IndentedJSON(http.StatusCreated, created)
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
			c.IndentedJSON(http.StatusInternalServerError, "Internal server error")
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
