package main

import (
	"html/template"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("no .env file found")
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
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
