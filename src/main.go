package main

import (
	"GoRedis/src/handlers"
	"GoRedis/src/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	// * Prepare the environment variables.
	// * For now we will just use hard coded environment key value pairs.
	utils.Environment()

	// * Get all the environment variables.
	port := os.Getenv("PORT")

	// * Logs the lived routes.
	log.Print(utils.TextGreen("Server live at port http://localhost:" + port))

	// * These are all the server routes.
	http.HandleFunc("/", handlers.ReadyHandler)
	http.HandleFunc("/redis", handlers.RedisHandler)

	// * You can setup the server here.
	s := &http.Server{
		Addr:           "localhost:7000",
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
