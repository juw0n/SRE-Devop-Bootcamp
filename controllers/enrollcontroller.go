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

type EnrollmentController struct {
	db  *db.Queries
	ctx context.Context
}

func NewEnrollmentController(db *db.Queries, ctx context.Context) *EnrollmentController {
	return &EnrollmentController{db, ctx}
}

// create enrollment handler
func (et *EnrollmentController) CreateEnrollment(ctx *gin.Context) {
	var payload *reqvalidate.CreateEnrollment

	// INFO: Log received request with payload details (including student and course IDs)
	log.Printf("INFO: Received POST request to create enrollment with data: %+v", payload)

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		log.Println("WARN: Failed to bind request payload:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	args := &db.CreateEnrollmentParams{
		EnrollmentDate: payload.EnrollmentDate,
		StudentID:      int32(payload.StudentID),
		CourseID:       int32(payload.CourseID),
	}

	// DEBUG: Log database call with parameters
	log.Printf("DEBUG: Calling CreateEnrollment with args: %+v", args)

	enrollment, err := et.db.CreateEnrollment(ctx, *args)

	if err != nil {
		log.Printf("ERROR: Error creating enrollment: %v", err)
		// Check if the error is a database error
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, sql.ErrConnDone) {
			// Use Internal Server Error for database-related issues
			log.Println("WARN: Database error occurred while creating enrollment.")
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Internal server error"})
		} else {
			// Use Bad Gateway or another appropriate status for other errors
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		}
		return
	}

	// INFO: Log successful enrollment creation
	log.Println("INFO: Enrollment created successfully. New enrollment ID:", enrollment.EnrollmentID)

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "enrollment": enrollment})
}

// Get all enrollment handler
func (et *EnrollmentController) GetAllEnrollments(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	// DEBUG: Log received request parameters
	log.Printf("DEBUG: Received GET request for all enrollments with page: %s and limit: %s", page, limit)

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

	args := &db.ListEnrollmentParams{
		Limit:  int32(intLimit),
		Offset: int32(offset),
	}

	// DEBUG: Log database call with parameters
	log.Printf("DEBUG: Calling ListEnrollment with args: %+v", args)

	enrollments, err := et.db.ListEnrollment(ctx, *args)
	if err != nil {
		log.Printf("ERROR: Error fetching enrollments: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if enrollments == nil {
		enrollments = []db.Enrollment{}
		log.Println("INFO: No enrollments found.")
	}

	// INFO: Log successful retrieval and number of enrollments found
	log.Printf("INFO: Retrieved %d enrollments successfully.", len(enrollments))

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(enrollments), "data": enrollments})
}

// Get a single enrollment handler
func (et *EnrollmentController) GetEnrollment(ctx *gin.Context) {
	enrollmentIDStr := ctx.Param("enrollmentID")

	// DEBUG: Log incoming request with enrollment ID
	log.Printf("DEBUG: Received GET request for enrollment with ID: %s", enrollmentIDStr)

	// Log the value of studentIDStr for debugging
	log.Println("Enrollment ID from URL:", enrollmentIDStr)

	// DEBUG: Log enrollment ID conversion
	log.Printf("DEBUG: Converting enrollment ID from string to int64: %s", enrollmentIDStr)
	// Convert studentID from string to int64
	enrollmentID, err := strconv.ParseInt(enrollmentIDStr, 10, 64)
	if err != nil {
		log.Println("WARN: Failed to parse enrollment ID:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid enrollment ID"})
		return
	}
	// Call the database method to get the student
	// DEBUG: Log database call with enrollment ID
	log.Printf("DEBUG: Calling GetEnrollment with enrollmentID: %d", enrollmentID)
	enrollment, err := et.db.GetEnrollment(ctx, enrollmentID)
	if err != nil {
		log.Printf("ERROR: Error fetching enrollment: %v", err)
		// Handle specific errors and return appropriate responses
		if err == sql.ErrNoRows {
			log.Println("WARN: Enrollment with ID not found:", enrollmentID)
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No enrollment with that ID exists"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Internal server error"})
		}
		return
	}
	// INFO: Log successful enrollment retrieval
	log.Println("INFO: Enrollment retrieved successfully with ID:", enrollment.EnrollmentID)

	// Return the student data as JSON
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "enrollment": enrollment})
}

// Delete a single enrollment handler
func (st *EnrollmentController) DeleteEnrollment(ctx *gin.Context) {
	enrollmentIDStr := ctx.Param("enrollmentID")

	// DEBUG: Log incoming request with enrollment ID
	log.Printf("DEBUG: Received DELETE request for enrollment with ID: %s", enrollmentIDStr)

	// Convert studentID from string to int64
	// DEBUG: Log enrollment ID conversion
	log.Printf("DEBUG: Converting enrollment ID from string to int64: %s", enrollmentIDStr)
	enrollmentID, err := strconv.ParseInt(enrollmentIDStr, 10, 64)
	if err != nil {
		log.Println("WARN: Failed to parse enrollment ID:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid enrollment ID"})
		return
	}

	// DEBUG: Log database call to check enrollment existence
	log.Printf("DEBUG: Calling GetEnrollment to check enrollment existence with enrollmentID: %d", enrollmentID)
	_, err = st.db.GetEnrollment(ctx, enrollmentID)
	if err != nil {
		log.Printf("ERROR: Error checking enrollment existence: %v", err)
		// Handle specific errors and return appropriate responses
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No enrollment with that ID exists"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		}
		return
	}

	// DEBUG: Log database call to delete enrollment
	log.Printf("DEBUG: Calling DeleteEnrollment with enrollmentID: %d", enrollmentID)
	err = st.db.DeleteEnrollment(ctx, enrollmentID)
	if err != nil {
		log.Printf("ERROR: Error deleting enrollment: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// INFO: Log successful enrollment deletion
	log.Println("INFO: Enrollment deleted successfully with ID:", enrollmentID)

	ctx.JSON(http.StatusNoContent, gin.H{"status": "success"})
}
