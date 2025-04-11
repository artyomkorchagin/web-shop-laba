package mssqlUser

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
	"errors"
	"fmt"
)

func (r *UserRepository) UpdateUser(ctx context.Context, u *types.User) error {
	// Validate input
	if u.Email == "" {
		return errors.New("invalid user or user ID")
	}

	// Prepare the update query dynamically
	query := "UPDATE users SET "
	params := []interface{}{}

	// Build the query based on the provided fields in the user object
	hasUpdates := false

	if u.DateOfBirth != "" {
		query += "[date_of_birth] = ?, "
		params = append(params, u.DateOfBirth)
		hasUpdates = true
	}
	if u.Username != "" {
		query += "[gender] = ?, "
		params = append(params, u.Username)
		hasUpdates = true
	}

	// Remove the trailing comma and add the WHERE clause
	if hasUpdates {
		query = query[:len(query)-2] + " WHERE [email] = ?"
		params = append(params, u.Email)
	} else {
		return errors.New("no fields provided for update")
	}

	// Execute the query
	result, err := r.db.ExecContext(ctx, query, params...)
	if err != nil {
		return fmt.Errorf("failed to execute update query: %w", err)
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
