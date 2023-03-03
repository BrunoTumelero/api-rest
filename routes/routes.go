package routes

import (
	"api-rest/controlers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controlers.TodosAlunos)
	r.GET("/:nome", controlers.Saudacao)
	r.POST("/alunos", controlers.CreateNew)
	r.GET("/alunos/:id", controlers.FindStudent)
	r.DELETE("/alunos/:id", controlers.DeleteStudent)
	r.Run()
}
