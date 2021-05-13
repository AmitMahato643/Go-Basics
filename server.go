package main

import (
	controller "gobasics/MasteringGoFollowUp/HandleJSON"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/postJson", controller.PostTask)
	http.ListenAndServe(":8080", router)
}
