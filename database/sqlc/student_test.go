package db

import (
	"context"
	"testing"

	"github.com/juw0n/SRE-Devop-Bootcamp/util"
	"github.com/stretchr/testify/require"
)

func TestCreateStudent(t *testing.T) {
	// Parse the date string into time.Time
	// dob, err := time.Parse("2006-01-02", "2000-04-02")
	// require.NoError(t, err)

	// Generate first and last names first
	firstName := util.RandomFirstName()
	lastName := util.RandomLastName()

	arg := CreateStudentParams{
		FirstName:    firstName,
		MiddleName:   util.RandomMiddleName(),
		LastName:     lastName,
		Gender:       util.RandomGender(),
		DateOfBirth:  util.RandomDateOfBirth(2000, 2099),
		PhoneNumber:  util.RandomPhoneNumber(),
		Email:        util.RandomEmail(firstName, lastName),
		YearOfEnroll: int32(util.RandomYearOfEnrollment(2005, 2099)),
		Country:      util.RandomCountries(),
		Major:        util.RandomMajor(),
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
