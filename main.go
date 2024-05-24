package main

import (
	"go-rest-api/models"
	"go-rest-api/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{ID: 1, Nome: "nome-1", Historia: "historia-1"},
		{ID: 2, Nome: "nome-2", Historia: "historia-2"},
	}

	routes.HandleRequest()
}
