package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-rest-api/models"
	"net/http"
	"strconv"
)

func RetornaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.ID) == id {
			json.NewEncoder(w).Encode(personalidade)
			return
		}
	}
}
