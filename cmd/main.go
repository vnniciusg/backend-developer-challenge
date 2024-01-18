package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/http"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/persistence"
)

// @title Backend Developer Challenger API Docs
// @version 0.0.1
// @description Backend Developer Challenger API Docs.
// @contact.name Vinnicius Santos
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @schemes http
func main() {
	conn, err := persistence.GetConnection()

	if err != nil {
		println(err.Error())
	}

	defer conn.Close()

	router := gin.Default()

	server := http.NewServer("8080", router, conn)

	server.Run()
}
