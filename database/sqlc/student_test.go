package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/juw0n/SRE-Devop-Bootcamp/util"
	"github.com/stretchr/testify/require"
)

// function to create a randomstudent entry
func createRandomStudent(t *testing.T) Student {
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
	require.Equal(t, arg.Major, student.Major)

	require.NotZero(t, student.StudentID)
	require.NotZero(t, student.CreatedAt)

	return student
}

// Unit test for createStudent operation
func TestCreateStudent(t *testing.T) {
	createRandomStudent(t)
}

// Unit test for GetStudent operation
func TestGetStudent(t *testing.T) {
	// create a random student entry
	student1 := createRandomStudent(t)
	student2, err := testQueries.GetStudent(context.Background(), student1.StudentID)

	require.NoError(t, err)
	require.NotEmpty(t, student2)

	require.Equal(t, student1.FirstName, student2.FirstName)
	require.Equal(t, student1.MiddleName, student2.MiddleName)
	require.Equal(t, student1.LastName, student2.LastName)
	require.Equal(t, student1.Gender, student2.Gender)
	require.True(t, student1.DateOfBirth.Equal(student2.DateOfBirth))
	require.Equal(t, student1.PhoneNumber, student2.PhoneNumber)
	require.Equal(t, student1.Email, student2.Email)
	require.Equal(t, student1.YearOfEnroll, student2.YearOfEnroll)
	require.Equal(t, student1.Country, student2.Country)
	require.Equal(t, student1.Major, student2.Major)

	require.NotZero(t, student2.StudentID)
	// Ensure StudentID is positive
	require.Greater(t, student2.StudentID, int64(0))
	require.NotZero(t, student2.CreatedAt)
	// Ensure CreatedAt is before or equal to the current time
	require.True(t, student2.CreatedAt.Before(time.Now().UTC()) || student2.CreatedAt.Equal(time.Now().UTC()))
}

// Unit test for updateStudent operation
func TestUpdateStudent(t *testing.T) {

	firstName := util.RandomFirstName()
	lastName := util.RandomLastName()
	// create a random student entry
	student1 := createRandomStudent(t)

	arg := UpdateStudentParams{
		StudentID:    student1.StudentID,
		FirstName:    firstName,
		MiddleName:   util.RandomMiddleName(),
		LastName:     lastName,
		Gender:       util.RandomGender(),
		DateOfBirth:  util.RandomDateOfBirth(3000, 3099),
		PhoneNumber:  util.RandomPhoneNumber(),
		Email:        util.RandomEmail(firstName, lastName),
		YearOfEnroll: int32(util.RandomYearOfEnrollment(3005, 3099)),
		Country:      util.RandomCountries(),
		Major:        util.RandomMajor(),
	}

	student2, err := testQueries.UpdateStudent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, student2)

	require.Equal(t, arg.FirstName, student2.FirstName)
	require.Equal(t, arg.MiddleName, student2.MiddleName)
	require.Equal(t, arg.LastName, student2.LastName)
	require.Equal(t, arg.Gender, student2.Gender)
	require.True(t, arg.DateOfBirth.Equal(student2.DateOfBirth))
	require.Equal(t, arg.PhoneNumber, student2.PhoneNumber)
	require.Equal(t, arg.Email, student2.Email)
	require.Equal(t, arg.YearOfEnroll, student2.YearOfEnroll)
	require.Equal(t, arg.Country, student2.Country)
	require.Equal(t, arg.Major, student2.Major)

	require.NotZero(t, student2.StudentID)
	// Ensure StudentID is positive
	require.Greater(t, student2.StudentID, int64(0))
	require.NotZero(t, student2.CreatedAt)
	// Ensure CreatedAt is before or equal to the current time
	require.True(t, student2.CreatedAt.Before(time.Now().UTC()) || student2.CreatedAt.Equal(time.Now().UTC()))
}

// Unit test for DeleteStudent operation
func TestDeleteStudent(t *testing.T) {
	// create a random student entry
	student1 := createRandomStudent(t)
	err := testQueries.DeleteStudent(context.Background(), student1.StudentID)
	require.NoError(t, err)

	student2, err := testQueries.GetStudent(context.Background(), student1.StudentID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, student2)
}

// Unit test for ListStudent operation
func TestListStudent(t *testing.T) {
	// Create multiple random students
	for i := 0; i < 10; i++ {
		createRandomStudent(t)
	}

	arg := ListStudentsParams{
		Limit:  5,
		Offset: 0,
	}

	// Retrieve the list of students
	students, err := testQueries.ListStudents(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, students, 5)

	for _, student := range students {
		require.NotEmpty(t, student.StudentID)
		require.NotEmpty(t, student.FirstName)
		require.NotEmpty(t, student.MiddleName)
		require.NotEmpty(t, student.LastName)
		require.NotEmpty(t, student.Gender)
		require.NotEmpty(t, student.DateOfBirth)
		require.NotEmpty(t, student.PhoneNumber)
		require.NotEmpty(t, student.Major)
	}
}
