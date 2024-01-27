package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/juw0n/SRE-Devop-Bootcamp/controllers"
)

type StudentRoutes struct {
	studentController controllers.StudentController
}

func NewStudentRoutes(studentController controllers.StudentController) StudentRoutes {
	return StudentRoutes{studentController}
}

func (sr *StudentRoutes) InitRoutes(rg *gin.RouterGroup) {
	router := rg.Group("students")
	router.POST("/", sr.studentController.CreateStudent)
	router.GET("/", sr.studentController.GetAllStudents)
	router.PATCH("/:studentId", sr.studentController.UpdateStudent)
	router.GET("/:studentId", sr.studentController.GetStudent)
	router.DELETE("/:studentId", sr.studentController.DeleteStudent)
}
