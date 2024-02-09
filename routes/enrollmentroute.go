package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/juw0n/SRE-Devop-Bootcamp/controllers"
)

type EnrollmentRoutes struct {
	enrollmentController controllers.EnrollmentController
}

func NewEnrollmentRoutes(enrollmentController controllers.EnrollmentController) EnrollmentRoutes {
	return EnrollmentRoutes{enrollmentController}
}

func (er *EnrollmentRoutes) InitRoutes(rg *gin.RouterGroup) {
	router := rg.Group("enrollments")
	router.POST("/", er.enrollmentController.CreateEnrollment)
	router.GET("/", er.enrollmentController.GetAllEnrollments)
	// router.PATCH("/:courseID", er.enrollmentController.UpdateCourse)
	router.GET("/:enrollmentID", er.enrollmentController.GetEnrollment)
	router.DELETE("/:enrollmentID", er.enrollmentController.DeleteEnrollment)
}
