package routes

import (
	"api-rest/controlers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/alunos", controlers.TodosAlunos)
	r.GET("/:nome", controlers.Saudacao)
	r.POST("/alunos", controlers.CreateNew)
	r.GET("/alunos/:id", controlers.FindStudent)
	r.DELETE("/alunos/:id", controlers.DeleteStudent)
	r.PATCH("/alunos/:id", controlers.EditStudent)
	r.GET("/alunos/cpf/:cpf", controlers.FindCPF)
	r.GET("/index", controlers.PagIndex)
	r.Run()
}
