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

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	args := &db.CreateEnrollmentParams{
		EnrollmentDate: payload.EnrollmentDate,
		StudentID:      int32(payload.StudentID),
		CourseID:       int32(payload.CourseID),
	}

	enrollment, err := et.db.CreateEnrollment(ctx, *args)

	if err != nil {
		log.Printf("Error creating enrollment: %v", err)
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
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "enrollment": enrollment})
}

// Get all enrollment handler
func (et *EnrollmentController) GetAllEnrollments(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, err := strconv.Atoi(page)
	if err != nil || intPage < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid page number"})
		return
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil || intLimit < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid limit value"})
		return
	}

	offset := (intPage - 1) * intLimit

	args := &db.ListEnrollmentParams{
		Limit:  int32(intLimit),
		Offset: int32(offset),
	}

	enrollments, err := et.db.ListEnrollment(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if enrollments == nil {
		enrollments = []db.Enrollment{}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(enrollments), "data": enrollments})
}

// Get a single enrollment handler
func (et *EnrollmentController) GetEnrollment(ctx *gin.Context) {
	enrollmentIDStr := ctx.Param("enrollmentID")

	// Log the value of studentIDStr for debugging
	log.Println("Enrollment ID from URL:", enrollmentIDStr)

	// Convert studentID from string to int64
	enrollmentID, err := strconv.ParseInt(enrollmentIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid enrollment ID"})
		return
	}
	// Call the database method to get the student
	enrollment, err := et.db.GetEnrollment(ctx, enrollmentID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No enrollment with that ID exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Internal server error"})
		return
	}
	// Return the student data as JSON
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "enrollment": enrollment})
}

// Delete a single enrollment handler
func (st *EnrollmentController) DeleteEnrollment(ctx *gin.Context) {
	enrollmentIDStr := ctx.Param("enrollmentID")

	// Convert studentID from string to int64
	enrollmentID, err := strconv.ParseInt(enrollmentIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid enrollment ID"})
		return
	}

	_, err = st.db.GetEnrollment(ctx, enrollmentID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No enrollment with that ID exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	err = st.db.DeleteEnrollment(ctx, enrollmentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"status": "success"})
}
