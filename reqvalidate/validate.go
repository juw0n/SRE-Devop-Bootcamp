package reqvalidate

import "time"

// Request Validation Structs
type CreateStudent struct {
	FirstName    string    `json:"first_name" binding:"required"`
	MiddleName   string    `json:"middle_name" binding:"required"`
	LastName     string    `json:"last_name" binding:"required"`
	Gender       string    `json:"gender" binding:"required"`
	DateOfBirth  time.Time `json:"date_of_birth" binding:"required"`
	PhoneNumber  string    `json:"phone_number" binding:"required"`
	Email        string    `json:"email" binding:"required"`
	YearOfEnroll int32     `json:"year_of_enroll" binding:"required"`
	Country      string    `json:"country" binding:"required"`
	Major        string    `json:"major" binding:"required"`
}

type UpdateStudent struct {
	StudentID    int64     `json:"student_id" binding:"required"`
	FirstName    string    `json:"first_name" binding:"required"`
	MiddleName   string    `json:"middle_name" binding:"required"`
	LastName     string    `json:"last_name" binding:"required"`
	Gender       string    `json:"gender" binding:"required, oneof=M F"`
	DateOfBirth  time.Time `json:"date_of_birth" binding:"required"`
	PhoneNumber  string    `json:"phone_number" binding:"required"`
	Email        string    `json:"email" binding:"required"`
	YearOfEnroll int32     `json:"year_of_enroll" binding:"required"`
	Country      string    `json:"country" binding:"required"`
	Major        string    `json:"major" binding:"required"`
}

type CreateCourse struct {
	CourseName string `json:"course_name" binding:"required"`
	Instructor string `json:"instructor" binding:"required"`
}

type UpdateCourse struct {
	CourseID   int64  `json:"course_id" binding:"required"`
	CourseName string `json:"course_name" binding:"required"`
	Instructor string `json:"instructor" binding:"required"`
}

type CreateEnrollment struct {
	EnrollmentDate time.Time `json:"enrollment_date" binding:"required"`
	StudentID      int64     `json:"student_id" binding:"required"`
	CourseID       int64     `json:"course_id" binding:"required"`
}
