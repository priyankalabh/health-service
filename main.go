package main

import (
	"flag"
	"health-service/router"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", ":7070", "this is a REST server port")
}

func main() {
	flag.Parse()
	server := gin.Default()
	router.AttachRouter(server)
	err := server.Run(port)
	if err != nil {
		log.Fatalf("Run error: %v", err)
	}
}
