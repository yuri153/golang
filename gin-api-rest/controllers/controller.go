package controller

import (
	"golang/api-go-rest/models"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	c.JSON(200, models.Students)
}

func Welcome(c *gin.Context) {
	name := c.Param("name")

	c.JSON(200, gin.H{
		"API says": "E ai " + name + " tudo beleza?",
	})
}
