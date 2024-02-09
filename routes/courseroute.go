package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/juw0n/SRE-Devop-Bootcamp/controllers"
)

type CourseRoutes struct {
	courseController controllers.CourseController
}

func NewCourseRoutes(courseController controllers.CourseController) CourseRoutes {
	return CourseRoutes{courseController}
}

func (cr *CourseRoutes) InitRoutes(rg *gin.RouterGroup) {
	router := rg.Group("courses")
	router.POST("/", cr.courseController.CreateCourse)
	router.GET("/", cr.courseController.GetAllCourses)
	router.PATCH("/:courseID", cr.courseController.UpdateCourse)
	router.GET("/:courseID", cr.courseController.GetCourse)
	router.DELETE("/:courseID", cr.courseController.DeleteCourse)
}
