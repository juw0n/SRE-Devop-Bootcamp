-- name: CreateCourse :one
INSERT INTO courses (
    course_name,
    instructor
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetCourse :one
SELECT * FROM courses
WHERE course_id = $1 LIMIT 1;

-- name: ListCourses :many
SELECT * FROM courses
ORDER BY course_id
LIMIT $1
OFFSET $2;

-- name: UpdateCourse :one
UPDATE courses
  set 
  course_name = $2,
  instructor = $3
WHERE course_id = $1
RETURNING *;

-- name: DeleteCourse :exec
DELETE FROM courses
WHERE course_id = $1;