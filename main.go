package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("no .env file found")
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World!!!</h1>"))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+getPort(), mux)
}
