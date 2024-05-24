package routes

import (
	"go-rest-api/controller"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/api/personalidades", controller.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
