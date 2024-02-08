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
	db  *db.Queries
	ctx context.Context
}

func NewStudentController(db *db.Queries, ctx context.Context) *StudentController {
	return &StudentController{db, ctx}
}

// create student handler
func (st *StudentController) CreateStudent(ctx *gin.Context) {
	var payload *reqvalidate.CreateStudent

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
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "student": student})
}

// Update student handler
func (st *StudentController) UpdateStudent(ctx *gin.Context) {
	var payload *reqvalidate.UpdateStudent
	// studentIDstr := ctx.Param("studentID")

	// // Convert studentID from string to int64
	// studentID, err := strconv.ParseInt(studentIDstr, 10, 64)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid student ID"})
	// 	return
	// }

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

	student, err := st.db.UpdateStudent(ctx, *args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that ID exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "student": student})
}

// Get a single student handler
func (st *StudentController) GetStudent(ctx *gin.Context) {
	studentIDStr := ctx.Param("studentID")

	// Log the value of studentIDStr for debugging
	log.Println("Student ID from URL:", studentIDStr)

	// Convert studentID from string to int64
	studentID, err := strconv.ParseInt(studentIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid student ID"})
		return
	}
	// Call the database method to get the student
	student, err := st.db.GetStudent(ctx, studentID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No student with that ID exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Internal server error"})
		return
	}
	// Return the student data as JSON
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "student": student})
}

// Delete a single student handler
func (st *StudentController) DeleteStudent(ctx *gin.Context) {
	studentIDStr := ctx.Param("studentID")

	// Convert studentID from string to int64
	studentID, err := strconv.ParseInt(studentIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid student ID"})
		return
	}

	_, err = st.db.GetStudent(ctx, studentID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No student with that ID exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	err = st.db.DeleteStudent(ctx, studentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"status": "success"})
}

// Get all students handler
func (st *StudentController) GetAllStudents(ctx *gin.Context) {
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

	args := &db.ListStudentsParams{
		Limit:  int32(intLimit),
		Offset: int32(offset),
	}

	students, err := st.db.ListStudents(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if students == nil {
		students = []db.Student{}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(students), "data": students})
}
