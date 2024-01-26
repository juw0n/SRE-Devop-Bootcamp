package controllers

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/juw0n/SRE-Devop-Bootcamp/database/sqlc"
	"github.com/juw0n/SRE-Devop-Bootcamp/schema"
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
	var payload *schema.CreateStudent

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
	var payload *schema.UpdateStudent
	studentID := ctx.Param("studentID")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	args := &db.UpdateStudentParams{
		StudentID:    studentID,
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
