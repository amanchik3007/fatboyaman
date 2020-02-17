package main

import (
	"bitrixapi/database"
	"bitrixapi/route"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	database.ConnectMySQL()
	defer database.CloseMySQL()

	port := os.Getenv("server_port")

	http.Handle("/", route.Handlers())

	fmt.Printf("Server running on port %s", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}