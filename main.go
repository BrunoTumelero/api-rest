package main

import (
	"api-rest/models"
	"api-rest/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Alfredo", CPF: "12345678920", RG: "14785236978"},
		{Nome: "Ana", CPF: "98745632111", RG: "12365478974"},
	}
	routes.HandleRequests()
}
