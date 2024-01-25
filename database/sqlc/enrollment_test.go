package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// createRandomEnrollment create a random enrollment into the database.
func createRandomEnrollment(t *testing.T) Enrollment {
	student := createRandomStudent(t)
	course := createRandomCourse(t)

	// Convert CourseID and studentID to int32
	courseID32 := int32(course.CourseID)
	studentID32 := int32(student.StudentID)

	arg := CreateEnrollmentParams{
		StudentID:      studentID32,
		CourseID:       courseID32,
		EnrollmentDate: time.Now(), // Or any other logic to set the date
	}

	enrollment, err := testQueries.CreateEnrollment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, enrollment)

	return enrollment
}

// Unit test for createEnrollment operation
func TestCreateEnrollment(t *testing.T) {
	// Create a random student
	student := createRandomStudent(t)
	require.NotEmpty(t, student)

	// Create a random course
	course := createRandomCourse(t)
	require.NotEmpty(t, course)

	// Convert CourseID and studentID to int32
	courseID32 := int32(course.CourseID)
	studentID32 := int32(student.StudentID)

	// Set up the CreateEnrollmentParams
	arg := CreateEnrollmentParams{
		StudentID: studentID32,
		CourseID:  courseID32,
		// Assume EnrollmentDate is the current date
		EnrollmentDate: time.Now(),
	}

	// Create the enrollment
	enrollment, err := testQueries.CreateEnrollment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, enrollment)

	// Validate the fields of the created enrollment
	require.Equal(t, arg.StudentID, enrollment.StudentID)
	require.Equal(t, arg.CourseID, enrollment.CourseID)
	// Assuming the date is stored without time component or with UTC timezone
	require.True(t, arg.EnrollmentDate.Truncate(24*time.Hour).Equal(enrollment.EnrollmentDate.Truncate(24*time.Hour)))
	require.NotZero(t, enrollment.EnrollmentID)
	require.NotZero(t, enrollment.CreatedAt)
}

// Unit test for getEnrollment operation
func TestGetEnrollment(t *testing.T) {
	// Create a random enrollment
	enrollment1 := createRandomEnrollment(t)

	// Retrieve the enrollment
	enrollment2, err := testQueries.GetEnrollment(context.Background(), enrollment1.EnrollmentID)
	require.NoError(t, err)
	require.NotEmpty(t, enrollment2)

	// Compare the retrieved enrollment with the created one
	require.Equal(t, enrollment1.EnrollmentID, enrollment2.EnrollmentID)
	require.Equal(t, enrollment1.StudentID, enrollment2.StudentID)
	require.Equal(t, enrollment1.CourseID, enrollment2.CourseID)
	require.WithinDuration(t, enrollment1.EnrollmentDate, enrollment2.EnrollmentDate, time.Second)
	require.True(t, enrollment1.CreatedAt.Before(time.Now().UTC()) || enrollment1.CreatedAt.Equal(time.Now().UTC()))
}

// Unit test for listEnrollment operation
func TestListEnrollment(t *testing.T) {
	// Create multiple random enrollments
	for i := 0; i < 10; i++ {
		createRandomEnrollment(t)
	}

	arg := ListEnrollmentParams{
		Limit:  5,
		Offset: 0,
	}

	// Retrieve the list of enrollments
	enrollments, err := testQueries.ListEnrollment(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, enrollments, 5)

	for _, enrollment := range enrollments {
		require.NotEmpty(t, enrollment.EnrollmentID)
		require.NotEmpty(t, enrollment.StudentID)
		require.NotEmpty(t, enrollment.CourseID)
		require.True(t, enrollment.CreatedAt.Before(time.Now().UTC()) || enrollment.CreatedAt.Equal(time.Now().UTC()))
	}
}

// Unit test for DeleteEnrollment operation
func TestDeleteEnrollment(t *testing.T) {
	// create a ramdon course
	enrollment1 := createRandomEnrollment(t)

	err := testQueries.DeleteEnrollment(context.Background(), enrollment1.EnrollmentID)
	require.NoError(t, err)

	course2, err := testQueries.GetEnrollment(context.Background(), enrollment1.EnrollmentID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, course2)
}
