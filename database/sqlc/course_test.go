package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/juw0n/SRE-Devop-Bootcamp/util"
	"github.com/stretchr/testify/require"
)

// function to create a randomCourse entry
func createRandomCourse(t *testing.T) Course {

	arg := CreateCourseParams{
		CourseName: util.RandomCourseName(),
		Instructor: util.RandomInstructorName(),
	}

	course, err := testQueries.CreateCourse(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, course)

	return course
}

// Unit test for createCourse operation
func TestCreateCourse(t *testing.T) {
	createRandomCourse(t)
}

// Unit test for GetCourse operation
func TestGetCourse(t *testing.T) {
	course1 := createRandomCourse(t)
	course2, err := testQueries.GetCourse(context.Background(), course1.CourseID)

	require.NoError(t, err)
	require.NotEmpty(t, course2)

	require.Equal(t, course1.CourseName, course2.CourseName)
	require.Equal(t, course1.Instructor, course2.Instructor)

	require.NotZero(t, course2.CourseID)
	// Ensure CourseID is positive
	require.Greater(t, course2.CourseID, int64(0))
	require.NotZero(t, course2.CreatedAt)
	// Ensure CreatedAt is before or equal to the current time
	require.True(t, course2.CreatedAt.Before(time.Now().UTC()) || course2.CreatedAt.Equal(time.Now().UTC()))
}

// Unit test for updateCourse operation
func TestUpateCourse(t *testing.T) {
	// create a ramdon course
	course1 := createRandomCourse(t)

	arg := UpdateCourseParams{
		CourseID:   course1.CourseID,
		CourseName: util.RandomCourseName(),
		Instructor: util.RandomInstructorName(),
	}

	course2, err := testQueries.UpdateCourse(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, course2)

	require.Equal(t, arg.CourseName, course2.CourseName)
	require.Equal(t, arg.Instructor, course2.Instructor)

	require.NotZero(t, course2.CourseID)
	// Ensure CourseID is positive
	require.Greater(t, course2.CourseID, int64(0))
	require.NotZero(t, course2.CreatedAt)
	// Ensure CreatedAt is before or equal to the current time
	require.True(t, course1.CreatedAt.Before(time.Now().UTC()) || course2.CreatedAt.Equal(time.Now().UTC()))

}

// Unit test for DeleteCourse operation
func TestDeleteCourse(t *testing.T) {
	// create a ramdon course
	course1 := createRandomCourse(t)

	err := testQueries.DeleteCourse(context.Background(), course1.CourseID)
	require.NoError(t, err)

	course2, err := testQueries.GetCourse(context.Background(), course1.CourseID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, course2)
}

// Unit test for ListCourse operation
func TestListCourse(t *testing.T) {
	// Create multiple random courses
	for i := 0; i < 10; i++ {
		createRandomCourse(t)
	}

	arg := ListCoursesParams{
		Limit:  5,
		Offset: 0,
	}

	// Retrieve the list of students
	courses, err := testQueries.ListCourses(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, courses, 5)

	for _, course := range courses {
		require.NotEmpty(t, course.CourseID)
		require.NotEmpty(t, course.Instructor)
		require.NotZero(t, course.CreatedAt)
	}
}
