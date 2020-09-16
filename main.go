package main

import (
	"./controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/logout", controllers.Logout).Methods("POST")
	r.HandleFunc("/del", controllers.DelAll).Methods("POST")
	r.HandleFunc("/refresh", controllers.Refresh).Methods("POST")

	http.ListenAndServe(":3000", r)
}
