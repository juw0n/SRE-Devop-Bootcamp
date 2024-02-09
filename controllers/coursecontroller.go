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

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	args := &db.CreateCourseParams{
		CourseName: payload.CourseName,
		Instructor: payload.Instructor,
	}

	course, err := cs.db.CreateCourse(ctx, *args)

	if err != nil {
		log.Printf("Error creating course: %v", err)
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
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "course": course})
}

// Update course handler
func (cs *CourseController) UpdateCourse(ctx *gin.Context) {
	var payload *reqvalidate.UpdateCourse

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	args := &db.UpdateCourseParams{
		CourseID:   payload.CourseID,
		CourseName: payload.CourseName,
		Instructor: payload.Instructor,
	}

	course, err := cs.db.UpdateCourse(ctx, *args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No course with that ID exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "course": course})
}

// Get a single course handler
func (cs *CourseController) GetCourse(ctx *gin.Context) {
	courseIDStr := ctx.Param("courseID")

	// Log the value of courseIDStr for debugging
	log.Println("Student ID from URL:", courseIDStr)

	// Convert courseID from string to int64
	courseID, err := strconv.ParseInt(courseIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid course ID"})
		return
	}

	// Call the database method to get the course
	course, err := cs.db.GetCourse(ctx, courseID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No course with that ID exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Internal server error"})
		return
	}
	// Return the student data as JSON
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "course": course})
}

// Get all courses handler
func (cs *CourseController) GetAllCourses(ctx *gin.Context) {
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

	args := &db.ListCoursesParams{
		Limit:  int32(intLimit),
		Offset: int32(offset),
	}

	courses, err := cs.db.ListCourses(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if courses == nil {
		courses = []db.Course{}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(courses), "data": courses})
}

// Delete a single course handler
func (cs *CourseController) DeleteCourse(ctx *gin.Context) {
	courseIDStr := ctx.Param("courseID")

	// Convert courseID from string to int64
	courseID, err := strconv.ParseInt(courseIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid course ID"})
		return
	}

	_, err = cs.db.GetCourse(ctx, courseID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No course with that ID exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	err = cs.db.DeleteCourse(ctx, courseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"status": "success"})
}
