package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/juw0n/SRE-Devop-Bootcamp/mocks"
	"github.com/juw0n/SRE-Devop-Bootcamp/reqvalidate"
	"github.com/juw0n/SRE-Devop-Bootcamp/util"
	// "go.uber.org/mock/gomock"
)

func TestCreateStudent(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Create a mock Querier
	mockDB := mocks.NewMockQuerier(mockCtrl)

	// Set up mock expectations
	// expectedStudent := util.GenerateStudent() // Define expected student data for mock expectation
	mockDB.EXPECT().
		CreateStudent(gomock.Any(), gomock.Any()).
		// Return(expectedStudent, nil).
		Times(1)

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

	ctx := context.Background()
	studentController := NewStudentController(mockDB, ctx)

	// Perform the HTTP request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		t.Fatal("Failed to marshal payload:", err)
	}
	reader := bytes.NewReader(payloadBytes)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", reader)
	studentController.CreateStudent(c)

	// Assert the response status code
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}
}
