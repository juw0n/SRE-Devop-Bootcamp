-- name: CreateStudent :one
INSERT INTO students (
  first_name,
  middle_name,
  last_name,
  gender,
  date_of_birth,
  phone_number,
  email,
  year_of_enroll,
  country,
  major
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: GetStudent :one
SELECT * FROM students
WHERE student_id = $1 LIMIT 1;

-- name: ListStudents :many
SELECT * FROM students
ORDER BY student_id
LIMIT $1
OFFSET $2;

-- name: UpdateStudent :one
UPDATE students
  set 
  first_name = $2,
  middle_name = $3,
  last_name = $4,
  gender = $5,
  date_of_birth = $6,
  phone_number = $7,
  email = $8,
  year_of_enroll = $9,
  country = $10,
  major = $11
WHERE student_id = $1
RETURNING *;

-- name: DeleteStudent :exec
DELETE FROM students
WHERE student_id = $1;