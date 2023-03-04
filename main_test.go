package main

import (
	"api-rest/controlers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRoutes() *gin.Engine {
	rotas := gin.Default()
	return rotas
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
