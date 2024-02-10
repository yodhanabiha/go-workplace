package controllers

import (
	"nabiha/project-golang/app/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestController struct {
	Config *config.Config
}

func NewTestController(cfg *config.Config) {
	cfg.Router.GET("/", func(gin *gin.Context) {
		gin.IndentedJSON(http.StatusOK, "Welcome to project Golang Nabiha!!")
	})

}
