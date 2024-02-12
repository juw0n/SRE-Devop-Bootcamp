package controllers

import (
	"context"

	db "github.com/juw0n/SRE-Devop-Bootcamp/database/sqlc"
	"github.com/stretchr/testify/mock"
)

// Database interface
type DBInterface interface {
	CreateStudent(ctx context.Context, args db.CreateStudentParams) (db.Student, error)
}

// MockDB is a mock implementation of the database interface for testing
type MockDB struct {
	mock.Mock
}
