package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateStudent(t *testing.T) {
	arg := CreateStudentParams{
		FirstName:    "Esther",
		MiddleName:   "Anuoluwa",
		LastName:     "Oluwagbenga",
		Gender:       "F",
		DateOfBirth:  "02042000",
		PhoneNumber:  "080678565433",
		Email:        "oduns07@gmail.com",
		YearOfEnroll: 2022,
		Country:      "Nigeria",
		Major:        "Marketing",
	}

	student, err := testQueries.CreateStudent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, student)

	require.NotZero(t, student.StudentID)
	require.NotZero(t, student.CreatedAt)

}
