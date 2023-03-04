package main

import (
	"api-rest/controlers"
	"api-rest/database"
	"api-rest/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CreatStudentMock() {
	student := models.Aluno{Nome: "Aluno teste", CPF: "12345678910",
		RG: "123654789"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Aluno
	database.DB.Delete(&student, ID)
}

func TestVerificationStatus(t *testing.T) {
	r := SetupRoutes()
	r.GET("/:nome", controlers.Saudacao)
	req, _ := http.NewRequest("GET", "/gui", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Resposta personalizada")
	mockResposta := `{"API diz": "Seja bem vindo " + nome}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockResposta, string(respostaBody))
}

func TestListstudents(t *testing.T) {
	database.ConectaComBancoDeDados()
	CreatStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutes()
	r.GET("/alunos", controlers.TodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestCPF(t *testing.T) {
	database.ConectaComBancoDeDados()
	CreatStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutes()
	r.GET("/alunos/cpf/:cpf", controlers.FindCPF)
	req, _ := http.NewRequest("GET", "alunos/cpf/numero do cpf", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
