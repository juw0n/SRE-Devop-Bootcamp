-- name: CreateEnrollment :one
INSERT INTO enrollments (
    enrollment_date, 
    student_id, 
    course_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetEnrollment :one
SELECT * FROM enrollments
WHERE enrollment_id = $1 LIMIT 1;

-- name: ListEnrollment :many
SELECT * FROM enrollments
ORDER BY enrollment_id
LIMIT $1
OFFSET $2;

-- name: DeleteEnrollment :exec
DELETE FROM enrollments
WHERE enrollment_id = $1;