package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provide all function to create db queries and transaction
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore create a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx execute a transaction within a databse transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	// Begin a new transaction
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// Create a new instance of Queries using the transaction
	q := New(tx)
	// Execute the provided function, passing it the Queries instance
	err = fn(q)
	// If an error occurred, rollback the transaction
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			// If rollback fails, return an error including both errors
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	// If no error occurred, commit the transaction
	return tx.Commit()
}

// UpdateStudentName updates the first name, middle name, and last name of a student.
func (store *Store) UpdateStudentName(ctx context.Context, studentID int, firstName, middleName, lastName string) error {
	// Define the SQL statement to update the student's name
	query := `
        UPDATE students
        SET first_name = $2, middle_name = $3, last_name = $4
        WHERE student_id = $1
    `

	// Execute the SQL statement within a transaction
	_, err := store.db.ExecContext(ctx, query, studentID, firstName, middleName, lastName)
	if err != nil {
		return fmt.Errorf("failed to update student name: %v", err)
	}
	return nil
}
