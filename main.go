package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/juw0n/SRE-Devop-Bootcamp/config"
	"github.com/juw0n/SRE-Devop-Bootcamp/controllers"
	dbConn "github.com/juw0n/SRE-Devop-Bootcamp/database/sqlc"
	"github.com/juw0n/SRE-Devop-Bootcamp/routes"
	_ "github.com/lib/pq"
)

var (
	server *gin.Engine
	db     *dbConn.Queries
	ctx    context.Context

	StudentController controllers.StudentController
	StudentRoutes     routes.StudentRoutes

	CourseController controllers.CourseController
	CourseRoutes     routes.CourseRoutes

	EnrollmentController controllers.EnrollmentController
	EnrollmentRoutes     routes.EnrollmentRoutes
)

func init() {
	ctx = context.TODO()
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("could not connect to postgres database: %v", err)
	}

	db = dbConn.New(conn)

	fmt.Println("PostgreSQL connected successfully...")

	StudentController = *controllers.NewStudentController(db, ctx)
	StudentRoutes = routes.NewStudentRoutes(StudentController)

	CourseController = *controllers.NewCourseController(db, ctx)
	CourseRoutes = routes.NewCourseRoutes(CourseController)

	EnrollmentController = *controllers.NewEnrollmentController(db, ctx)
	EnrollmentRoutes = routes.NewEnrollmentRoutes(EnrollmentController)

	server = gin.Default()
}

func main() {
	// Load configuration
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Set up CORS
	corsConfig := cors.DefaultConfig()
	if config.Origin != "" {
		corsConfig.AllowOrigins = []string{config.Origin}
	} else {
		// Use a default value or a wildcard (not recommended for production)
		corsConfig.AllowOrigins = []string{"*"}
	}
	corsConfig.AllowCredentials = true

	// Apply CORS middleware to your server
	server.Use(cors.New(corsConfig))

	// Set up routing
	router := server.Group("/api/v1")

	router.GET("/healthchecker", func(ctx *gin.Context) {
		// Log successful health check
		log.Println("INFO: Health check passed successfully.")
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Welcome to my simple REST API using Golang and PostgreSQL"})
	})
	// Initialize Student routes
	StudentRoutes.InitRoutes(router)

	// Initialize Course routes
	CourseRoutes.InitRoutes(router)

	// Initialize Enrollment routes
	EnrollmentRoutes.InitRoutes(router)

	// Handle no route found
	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": fmt.Sprintf("Route %s not found", ctx.Request.URL)})
	})
	// Start the server
	log.Fatal(server.Run(":" + config.ServerPort))
}
