package controllers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/juw0n/SRE-Devop-Bootcamp/mocks"
	"github.com/juw0n/SRE-Devop-Bootcamp/reqvalidate"
	"github.com/juw0n/SRE-Devop-Bootcamp/util"
	"go.uber.org/mock/gomock"
)

func TestCreateStudent(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocks.NewMockQuerier(mockCtrl)
	ctx := context.Background()

	studentController := NewStudentController(mockDB, ctx)

	// Prepare a sample request payload
	firstName := util.RandomFirstName()
	lastName := util.RandomLastName()

	payload := &reqvalidate.CreateStudent{
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

	// Set up mock expectations
	mockDB.EXPECT().
		CreateStudent(gomock.Any(), gomock.Any()).
		Return(expectedStudent, nil).
		Times(1)

	// Perform the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", payload)
	studentController.CreateStudent(c)

	// // Assert the response
	// assert.Equal(t, http.StatusBadGateway, w.Code)
	// // Add more assertions based on the expected behavior of your controller
}
