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

type StudentController struct {
	db  db.Querier
	ctx context.Context
}

// type StudentController struct {
// 	db  *db.Queries
// 	ctx context.Context
// }

func NewStudentController(db db.Querier, ctx context.Context) *StudentController {
	return &StudentController{db, ctx}
}

// create student handler
func (st *StudentController) CreateStudent(ctx *gin.Context) {
	var payload *reqvalidate.CreateStudent

	// INFO: Log incoming request with payload details
	log.Println("INFO: Received POST request to create student with data:", payload)

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	// now := time.Now()
	args := &db.CreateStudentParams{
		FirstName:    payload.FirstName,
		MiddleName:   payload.MiddleName,
		LastName:     payload.LastName,
		Gender:       payload.Gender,
		DateOfBirth:  payload.DateOfBirth,
		PhoneNumber:  payload.PhoneNumber,
		Email:        payload.Email,
		YearOfEnroll: payload.YearOfEnroll,
		Country:      payload.Country,
		Major:        payload.Major,
	}

	// DEBUG: Log database call with parameters
	log.Printf("DEBUG: Calling CreateStudent with args: %+v", args)

	student, err := st.db.CreateStudent(ctx, *args)

	if err != nil {
		log.Printf("Error creating student: %v", err)
		// Check if the error is a database error
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, sql.ErrConnDone) {
			// Use Internal Server Error for database-related issues
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Internal server error"})
		} else {
			// Use Bad Gateway or another appropriate status for other errors
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		}
		return
	}

	// INFO: Log successful student creation
	log.Println("INFO: Student created successfully. New student ID:", student.StudentID)

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "student": student})
}

// Update student handler
func (st *StudentController) UpdateStudent(ctx *gin.Context) {
	var payload *reqvalidate.UpdateStudent

	// INFO: Log incoming request with payload details
	log.Println("INFO: Received PUT request to update student with ID:", ctx.Param("studentID"), "and data:", payload)

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	args := &db.UpdateStudentParams{
		StudentID:    payload.StudentID,
		FirstName:    payload.FirstName,
		MiddleName:   payload.MiddleName,
		LastName:     payload.LastName,
		Gender:       payload.Gender,
		DateOfBirth:  payload.DateOfBirth,
		PhoneNumber:  payload.PhoneNumber,
		Email:        payload.Email,
		YearOfEnroll: payload.YearOfEnroll,
		Country:      payload.Country,
		Major:        payload.Major,
	}

	// DEBUG: Log database call with parameters
	log.Printf("DEBUG: Calling UpdateStudent with args: %+v", args)

	student, err := st.db.UpdateStudent(ctx, *args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No student with that ID exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// INFO: Log successful student update
	log.Println("INFO: Student updated successfully with ID:", student.StudentID)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "student": student})
}

// Get a single student handler
func (st *StudentController) GetStudent(ctx *gin.Context) {
	studentIDStr := ctx.Param("studentID")

	// DEBUG: Log incoming request with student ID
	log.Printf("DEBUG: Received GET request for student with ID: %s", studentIDStr)

	// Log the value of studentIDStr for debugging
	log.Println("Student ID from URL:", studentIDStr)

	// DEBUG: Log student ID conversion
	log.Printf("DEBUG: Converting student ID from string to int64: %s", studentIDStr)

	// Convert studentID from string to int64
	studentID, err := strconv.ParseInt(studentIDStr, 10, 64)
	if err != nil {
		log.Println("WARN: Failed to parse student ID:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid student ID"})
		return
	}

	// DEBUG: Log database call with student ID
	log.Printf("DEBUG: Calling GetStudent with studentID: %d", studentID)

	// Call the database method to get the student
	student, err := st.db.GetStudent(ctx, studentID)
	if err != nil {
		log.Printf("ERROR: Error fetching student: %v", err)
		if err == sql.ErrNoRows {
			log.Println("WARN: Student with ID not found:", studentID)
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No student with that ID exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Internal server error"})
		return
	}
	// INFO: Log successful student retrieval
	log.Println("INFO: Student retrieved successfully with ID:", student.StudentID)

	// Return the student data as JSON
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "student": student})
}

// Delete a single student handler
func (st *StudentController) DeleteStudent(ctx *gin.Context) {
	studentIDStr := ctx.Param("studentID")

	// DEBUG: Log incoming request with student ID
	log.Printf("DEBUG: Received DELETE request for student with ID: %s", studentIDStr)

	// DEBUG: Log student ID conversion
	log.Printf("DEBUG: Converting student ID from string to int64: %s", studentIDStr)
	// Convert studentID from string to int64
	studentID, err := strconv.ParseInt(studentIDStr, 10, 64)
	if err != nil {
		log.Println("WARN: Failed to parse student ID:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid student ID"})
		return
	}

	// DEBUG: Log database call to check student existence
	log.Printf("DEBUG: Calling GetStudent to check student existence with studentID: %d", studentID)
	_, err = st.db.GetStudent(ctx, studentID)
	if err != nil {
		log.Printf("ERROR: Error checking student existence: %v", err)
		if err == sql.ErrNoRows {
			log.Println("WARN: Student with ID not found before deletion:", studentID)
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No student with that ID exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// DEBUG: Log database call to delete student
	log.Printf("DEBUG: Calling DeleteStudent with studentID: %d", studentID)
	err = st.db.DeleteStudent(ctx, studentID)
	if err != nil {
		log.Printf("ERROR: Error deleting student: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// INFO: Log successful student deletion
	log.Println("INFO: Student deleted successfully with ID:", studentID)

	ctx.JSON(http.StatusNoContent, gin.H{"status": "success"})
}

// Get all students handler
func (st *StudentController) GetAllStudents(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	// DEBUG: Log received request parameters
	log.Printf("DEBUG: Received GET request for all students with page: %s and limit: %s", page, limit)

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

	args := &db.ListStudentsParams{
		Limit:  int32(intLimit),
		Offset: int32(offset),
	}

	// DEBUG: Log database call with parameters
	log.Printf("DEBUG: Calling ListStudents with args: %+v", args)

	students, err := st.db.ListStudents(ctx, *args)
	if err != nil {
		log.Printf("ERROR: Error fetching students: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if students == nil {
		students = []db.Student{}
		log.Println("INFO: No students found.")
	}
	// INFO: Log successful retrieval and number of students found
	log.Printf("INFO: Retrieved %d students successfully.", len(students))

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(students), "data": students})
}
