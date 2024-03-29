// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: course.sql

package db

import (
	"context"
)

const createCourse = `-- name: CreateCourse :one
INSERT INTO courses (
    course_name,
    instructor
) VALUES (
  $1, $2
)
RETURNING course_id, course_name, instructor, created_at
`

type CreateCourseParams struct {
	CourseName string `json:"course_name"`
	Instructor string `json:"instructor"`
}

func (q *Queries) CreateCourse(ctx context.Context, arg CreateCourseParams) (Course, error) {
	row := q.db.QueryRowContext(ctx, createCourse, arg.CourseName, arg.Instructor)
	var i Course
	err := row.Scan(
		&i.CourseID,
		&i.CourseName,
		&i.Instructor,
		&i.CreatedAt,
	)
	return i, err
}

const deleteCourse = `-- name: DeleteCourse :exec
DELETE FROM courses
WHERE course_id = $1
`

func (q *Queries) DeleteCourse(ctx context.Context, courseID int64) error {
	_, err := q.db.ExecContext(ctx, deleteCourse, courseID)
	return err
}

const getCourse = `-- name: GetCourse :one
SELECT course_id, course_name, instructor, created_at FROM courses
WHERE course_id = $1 LIMIT 1
`

func (q *Queries) GetCourse(ctx context.Context, courseID int64) (Course, error) {
	row := q.db.QueryRowContext(ctx, getCourse, courseID)
	var i Course
	err := row.Scan(
		&i.CourseID,
		&i.CourseName,
		&i.Instructor,
		&i.CreatedAt,
	)
	return i, err
}

const listCourses = `-- name: ListCourses :many
SELECT course_id, course_name, instructor, created_at FROM courses
ORDER BY course_id
LIMIT $1
OFFSET $2
`

type ListCoursesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCourses(ctx context.Context, arg ListCoursesParams) ([]Course, error) {
	rows, err := q.db.QueryContext(ctx, listCourses, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Course
	for rows.Next() {
		var i Course
		if err := rows.Scan(
			&i.CourseID,
			&i.CourseName,
			&i.Instructor,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCourse = `-- name: UpdateCourse :one
UPDATE courses
  set 
  course_name = $2,
  instructor = $3
WHERE course_id = $1
RETURNING course_id, course_name, instructor, created_at
`

type UpdateCourseParams struct {
	CourseID   int64  `json:"course_id"`
	CourseName string `json:"course_name"`
	Instructor string `json:"instructor"`
}

func (q *Queries) UpdateCourse(ctx context.Context, arg UpdateCourseParams) (Course, error) {
	row := q.db.QueryRowContext(ctx, updateCourse, arg.CourseID, arg.CourseName, arg.Instructor)
	var i Course
	err := row.Scan(
		&i.CourseID,
		&i.CourseName,
		&i.Instructor,
		&i.CreatedAt,
	)
	return i, err
}
