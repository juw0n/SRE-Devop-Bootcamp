package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateStudent(t *testing.T) {
	// Parse the date string into time.Time
	dob, err := time.Parse("2006-01-02", "2000-04-02")
	require.NoError(t, err)

	arg := CreateStudentParams{
		FirstName:    "Esther",
		MiddleName:   "Anuoluwa",
		LastName:     "Oluwagbenga",
		Gender:       "F",
		DateOfBirth:  dob,
		PhoneNumber:  "080678565433",
		Email:        "oduns07@gmail.com",
		YearOfEnroll: 2022,
		Country:      "Nigeria",
		Major:        "Marketing",
	}

	student, err := testQueries.CreateStudent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, student)

	require.Equal(t, arg.FirstName, student.FirstName)
	require.Equal(t, arg.MiddleName, student.MiddleName)
	require.Equal(t, arg.LastName, student.LastName)
	require.Equal(t, arg.Gender, student.Gender)
	require.True(t, arg.DateOfBirth.UTC().Equal(student.DateOfBirth.UTC()))
	require.Equal(t, arg.PhoneNumber, student.PhoneNumber)
	require.Equal(t, arg.Email, student.Email)
	require.Equal(t, arg.YearOfEnroll, student.YearOfEnroll)
	require.Equal(t, arg.Country, student.Country)

	require.NotZero(t, student.StudentID)
	require.NotZero(t, student.CreatedAt)

	require.Equal(t, arg.Email, student.Email)
	require.Equal(t, arg.Major, student.Major)

}
