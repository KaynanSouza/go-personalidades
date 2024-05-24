package controller

import (
	"encoding/json"
	"go-rest-api/models"
	"net/http"
)

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
