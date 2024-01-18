package main

import (
	"fmt"

	_ "github.com/lib/pq"

	_ "github.com/vnniciusg/backend-developer-challenge/docs"

	"github.com/gin-gonic/gin"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/http"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/persistence"
	"github.com/vnniciusg/backend-developer-challenge/internal/pkg/persistence/utils"
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
		fmt.Printf("Error connecting to the database: %s\n", err.Error())
		panic(err.Error())
	}

	defer conn.Close()

	utils.RunMigrations()

	router := gin.Default()

	server := http.NewServer("8080", router, conn)

	err = server.Run()

	if err != nil {
		fmt.Printf("Error running server: %s\n", err.Error())
		panic(err.Error())
	}
}
