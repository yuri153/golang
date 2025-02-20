package routes

import (
	controller "golang/api-go-rest/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/students", controller.GetStudents)
	r.GET("/students/:id", controller.GetStudentsById)
	r.GET("/students/cpf/:cpf", controller.GetStudentByCpf)
	r.DELETE("/students/:id", controller.DeleteStudent)
	r.PATCH("/students/:id", controller.UpdateStudent)
	r.POST("/students", controller.CreateStudent)
	r.GET("/:name", controller.Welcome)
	r.Run(":8080")
}
