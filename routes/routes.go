package routes

import (
	"github.com/gorilla/mux"
	"go-rest-api/controller"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Home)
	r.HandleFunc("/api/personalidades", controller.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controller.RetornaPersonalidade).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
