package controllers

import (
	"context"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	db "github.com/juw0n/SRE-Devop-Bootcamp/database/sqlc"
	dbMocks "github.com/juw0n/SRE-Devop-Bootcamp/mocks"
	// "go.uber.org/mock/gomock"
)

func TestCreateStudent(t *testing.T) {
	// Create a new instance of the mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock for the db.Queries interface
	mockDB := dbMocks.NewMockQuerier(ctrl)

	// Create a context
	ctx := context.TODO()

	// Create a new instance of the student controller with the mock
	stController := NewStudentController(mockDB, ctx)

	// Generate first and last names first
	// firstName := util.RandomFirstName()
	// lastName := util.RandomLastName()

	// // Prepare sample input payload
	// payload := &reqvalidate.CreateStudent{
	// 	FirstName:    firstName,
	// 	LastName:     lastName,
	// 	Gender:       util.RandomGender(),
	// 	DateOfBirth:  util.RandomDateOfBirth(2000, 2099),
	// 	PhoneNumber:  util.RandomPhoneNumber(),
	// 	Email:        util.RandomEmail(firstName, lastName),
	// 	YearOfEnroll: int32(util.RandomYearOfEnrollment(2005, 2099)),
	// 	Country:      util.RandomCountries(),
	// 	Major:        util.RandomMajor(),
	// }

	// Mock the behavior of CreateStudent method in the mock
	mockDB.EXPECT().CreateStudent(ctx, gomock.Any()).Return(&db.Student{}, nil)

	// Create a mock gin context
	ginCtx, _ := gin.CreateTestContext(nil)
	ginCtx.Set("testKey", "testValue")

	// Call the function under test
	stController.CreateStudent(ginCtx)

	// Check the response
	if ginCtx.Writer.Status() != http.StatusCreated {
		t.Errorf("Expected status %d but got %d", http.StatusCreated, ginCtx.Writer.Status())
	}
}
