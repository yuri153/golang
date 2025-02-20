package routes

import (
	controller "golang/api-go-rest/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/students", controller.GetStudents)
	r.GET("/:name", controller.Welcome)
	r.Run(":8080")
}
