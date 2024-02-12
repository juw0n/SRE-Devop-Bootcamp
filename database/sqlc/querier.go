// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	CreateCourse(ctx context.Context, arg CreateCourseParams) (Course, error)
	CreateEnrollment(ctx context.Context, arg CreateEnrollmentParams) (Enrollment, error)
	CreateStudent(ctx context.Context, arg CreateStudentParams) (Student, error)
	DeleteCourse(ctx context.Context, courseID int64) error
	DeleteEnrollment(ctx context.Context, enrollmentID int64) error
	DeleteStudent(ctx context.Context, studentID int64) error
	GetCourse(ctx context.Context, courseID int64) (Course, error)
	GetEnrollment(ctx context.Context, enrollmentID int64) (Enrollment, error)
	GetStudent(ctx context.Context, studentID int64) (Student, error)
	ListCourses(ctx context.Context, arg ListCoursesParams) ([]Course, error)
	ListEnrollment(ctx context.Context, arg ListEnrollmentParams) ([]Enrollment, error)
	ListStudents(ctx context.Context, arg ListStudentsParams) ([]Student, error)
	UpdateCourse(ctx context.Context, arg UpdateCourseParams) (Course, error)
	UpdateStudent(ctx context.Context, arg UpdateStudentParams) (Student, error)
}

var _ Querier = (*Queries)(nil)
