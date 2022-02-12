package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {


	app := &App{}
	app.Initialize()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println(fmt.Sprintf("http://localhost:%s", port))
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), app.Router)
	if err != nil {
		log.Fatal(fmt.Sprintf("Listen on %s error", port))
		return
	}

}
