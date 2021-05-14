package main

import (
	handler "gobasics/Handler"
	"gobasics/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	utils.LoadEnv()

	r := mux.NewRouter()
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		http.Redirect(rw, r, "/home", http.StatusPermanentRedirect)
	})
	r.HandleFunc("/home", handler.HomeHandler).Name("home")
	r.HandleFunc("/login", handler.LoginHandler)
	r.HandleFunc("/authenticatelogin", handler.HandleCallback)

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
