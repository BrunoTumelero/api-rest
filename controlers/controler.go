package controlers

import (
	"api-rest/database"
	"api-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TodosAlunos(c *gin.Context) {
	var students []models.Aluno
	database.DB.Find(&students)
	c.JSON(200, students)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz": "Seja bem vindo " + nome,
	})
}

func CreateNew(c *gin.Context) {
	var student models.Aluno
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func FindStudent(c *gin.Context) {
	var student models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Aluno
	id := c.Params.ByName("id")
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Aluno inexistente"})
	}
	database.DB.Delete(&student, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Aluno deletado com sucesso"})
}

func EditStudent(c *gin.Context) {
	var student models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}
