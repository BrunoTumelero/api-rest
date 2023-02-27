package controlers

import (
	"api-rest/models"

	"github.com/gin-gonic/gin"
)

func TodosAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz": "Seja bem vindo " + nome,
	})
}
