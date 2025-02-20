package controller

import (
	"golang/api-go-rest/database"
	"golang/api-go-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)

	c.JSON(http.StatusOK, students)
}

func Welcome(c *gin.Context) {
	name := c.Param("name")

	c.JSON(http.StatusOK, gin.H{
		"API says": "E ai " + name + " tudo beleza?",
	})
}

func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&student)

	c.JSON(http.StatusOK, student)
}

func GetStudentsById(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Not found": "Student not found!"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	database.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{"Deleted": "Student deleted!"})
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&student).UpdateColumns(&student)

	c.JSON(http.StatusOK, student)
}

func GetStudentByCpf(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")

	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Not found": "Student not found!"})
		return
	}

	c.JSON(http.StatusOK, student)
}
