package controllers

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/juw0n/SRE-Devop-Bootcamp/database/sqlc"
	"github.com/juw0n/SRE-Devop-Bootcamp/reqvalidate"
)

type CourseController struct {
	db  *db.Queries
	ctx context.Context
}

func NewCourseController(db *db.Queries, ctx context.Context) *CourseController {
	return &CourseController{db, ctx}
}

// create course handler
func (cs *CourseController) CreateCourse(ctx *gin.Context) {
	var payload *reqvalidate.CreateCourse

	// INFO: Log received request with payload details
	log.Println("INFO: Received POST request to create course with data:", payload)

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		log.Println("WARN: Failed to bind request payload:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	args := &db.CreateCourseParams{
		CourseName: payload.CourseName,
		Instructor: payload.Instructor,
	}

	// DEBUG: Log database call with parameters
	log.Printf("DEBUG: Calling CreateCourse with args: %+v", args)

	course, err := cs.db.CreateCourse(ctx, *args)

	if err != nil {
		log.Printf("Error creating course: %v", err)
		// Check if the error is a database error
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, sql.ErrConnDone) {
			// Use Internal Server Error for database-related issues
			log.Println("WARN: Database error occurred while creating course.")
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Internal server error"})
		} else {
			// Use Bad Gateway or another appropriate status for other errors
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		}
		return
	}
	// INFO: Log successful course creation
	log.Println("INFO: Course created successfully. New course ID:", course.CourseID)

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "course": course})
}

// Update course handler
func (cs *CourseController) UpdateCourse(ctx *gin.Context) {
	var payload *reqvalidate.UpdateCourse

	// INFO: Log received request with payload details (including course ID)
	log.Printf("INFO: Received PUT request to update course with ID: %s and data: %+v", ctx.Param("courseID"), payload)

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		log.Println("WARN: Failed to bind request payload:", err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	args := &db.UpdateCourseParams{
		CourseID:   payload.CourseID,
		CourseName: payload.CourseName,
		Instructor: payload.Instructor,
	}

	// DEBUG: Log database call with parameters
	log.Printf("DEBUG: Calling UpdateCourse with args: %+v", args)

	course, err := cs.db.UpdateCourse(ctx, *args)

	if err != nil {
		log.Printf("ERROR: Error updating course: %v", err)
		if err == sql.ErrNoRows {
			log.Println("WARN: Course with ID not found:", args.CourseID)
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No course with that ID exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}
	// INFO: Log successful course update
	log.Println("INFO: Course updated successfully with ID:", course.CourseID)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "course": course})
}

// Get a single course handler
func (cs *CourseController) GetCourse(ctx *gin.Context) {
	courseIDStr := ctx.Param("courseID")

	// DEBUG: Log incoming request with course ID
	log.Printf("DEBUG: Received GET request for course with ID: %s", courseIDStr)

	// Log the value of courseIDStr for debugging
	log.Println("Student ID from URL:", courseIDStr)

	// Convert courseID from string to int64
	log.Printf("DEBUG: Converting course ID from string to int64: %s", courseIDStr)
	courseID, err := strconv.ParseInt(courseIDStr, 10, 64)
	if err != nil {
		log.Println("WARN: Failed to parse course ID:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid course ID"})
		return
	}

	// DEBUG: Log database call with course ID
	log.Printf("DEBUG: Calling GetCourse with courseID: %d", courseID)
	// Call the database method to get the course
	course, err := cs.db.GetCourse(ctx, courseID)
	if err != nil {
		log.Printf("ERROR: Error fetching course: %v", err)
		if err == sql.ErrNoRows {
			log.Println("WARN: Course with ID not found:", courseID)
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No course with that ID exists"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Internal server error"})
		}
		return
	}

	// INFO: Log successful course retrieval
	log.Println("INFO: Course retrieved successfully with ID:", course.CourseID)
	// Return the student data as JSON
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "course": course})
}

// Get all courses handler
func (cs *CourseController) GetAllCourses(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	// DEBUG: Log received request parameters
	log.Printf("DEBUG: Received GET request for all courses with page: %s and limit: %s", page, limit)

	intPage, err := strconv.Atoi(page)
	if err != nil || intPage < 1 {
		log.Printf("WARN: Invalid page number: %s", page)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid page number"})
		return
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil || intLimit < 1 {
		log.Printf("WARN: Invalid limit value: %s", limit)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid limit value"})
		return
	}

	offset := (intPage - 1) * intLimit

	args := &db.ListCoursesParams{
		Limit:  int32(intLimit),
		Offset: int32(offset),
	}

	// DEBUG: Log database call with parameters
	log.Printf("DEBUG: Calling ListCourses with args: %+v", args)

	courses, err := cs.db.ListCourses(ctx, *args)
	if err != nil {
		log.Printf("ERROR: Error fetching courses: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// Handle case where no courses found
	if courses == nil {
		courses = []db.Course{}
		log.Println("INFO: No courses found.")
	}

	// INFO: Log successful retrieval and number of courses found
	log.Printf("INFO: Retrieved %d courses successfully.", len(courses))

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(courses), "data": courses})
}

// Delete a single course handler
func (cs *CourseController) DeleteCourse(ctx *gin.Context) {
	courseIDStr := ctx.Param("courseID")

	// DEBUG: Log incoming request with course ID
	log.Printf("DEBUG: Received DELETE request for course with ID: %s", courseIDStr)

	// DEBUG: Log course ID conversion
	log.Printf("DEBUG: Converting course ID from string to int64: %s", courseIDStr)

	// Convert courseID from string to int64
	courseID, err := strconv.ParseInt(courseIDStr, 10, 64)
	if err != nil {
		log.Println("WARN: Failed to parse course ID:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid course ID"})
		return
	}

	// DEBUG: Log database call to check course existence
	log.Printf("DEBUG: Calling GetCourse to check course existence with courseID: %d", courseID)
	_, err = cs.db.GetCourse(ctx, courseID)
	if err != nil {
		log.Printf("ERROR: Error checking course existence: %v", err)
		if err == sql.ErrNoRows {
			log.Println("WARN: Course with ID not found before deletion:", courseID)
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No course with that ID exists"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		}
		return
	}

	// DEBUG: Log database call to delete course
	log.Printf("DEBUG: Calling DeleteCourse with courseID: %d", courseID)
	err = cs.db.DeleteCourse(ctx, courseID)
	if err != nil {
		log.Printf("ERROR: Error deleting course: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// INFO: Log successful course deletion
	log.Println("INFO: Course deleted successfully with ID:", courseID)

	ctx.JSON(http.StatusNoContent, gin.H{"status": "success"})
}
