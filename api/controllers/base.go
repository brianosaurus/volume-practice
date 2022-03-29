package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brianosaurus/volume-practice/api/middlewares"
	"github.com/gin-gonic/gin"
)

//Server xxx
type Server struct {
	Router *gin.Engine
}

var errList = make(map[string]string)

//Initialize connects to various services
func (server *Server) Initialize() {
	fmt.Println("XXXXXXXXXXXX")

	// If you are using mysql, i added support for you here(dont forgot to edit the .env file)
	server.Router = gin.Default()
	server.Router.Use(middlewares.CORSMiddleware())

	server.initializeRoutes()
}

//Run start the system
func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
