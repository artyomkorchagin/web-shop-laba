package mssqlUser

import (
	"context"
	"errors"
	"fmt"
	"socialsecurity/internal/types"
)

func (r *mssqlUserRepository) UpdateUser(ctx context.Context, u *types.User) error {
	// Validate input
	if u.Email == "" {
		return errors.New("invalid user or user ID")
	}

	// Prepare the update query dynamically
	query := "UPDATE users SET "
	params := []interface{}{}

	// Build the query based on the provided fields in the user object
	hasUpdates := false

	if u.FirstName != "" {
		query += "[first_name] = ?, "
		params = append(params, u.FirstName)
		hasUpdates = true
	}
	if u.SecondName != "" {
		query += "[second_name] = ?, "
		params = append(params, u.SecondName)
		hasUpdates = true
	}
	if u.MiddleName != "" {
		query += "[middle_name] = ?, "
		params = append(params, u.MiddleName)
		hasUpdates = true
	}
	if u.DateOfBirth != "" {
		query += "[date_of_birth] = ?, "
		params = append(params, u.DateOfBirth)
		hasUpdates = true
	}
	if u.Gender != "" {
		query += "[gender] = ?, "
		params = append(params, u.Gender)
		hasUpdates = true
	}
	if u.Address != "" {
		query += "[address] = ?, "
		params = append(params, u.Address)
		hasUpdates = true
	}
	if u.PhoneNumber != "" {
		query += "[phone_number] = ?, "
		params = append(params, u.PhoneNumber)
		hasUpdates = true
	}
	if u.SNILS != "" {
		query += "[snils] = ?, "
		params = append(params, u.SNILS)
		hasUpdates = true
	}
	if u.PassportSeries != "" {
		query += "[passport_series] = ?, "
		params = append(params, u.PassportSeries)
		hasUpdates = true
	}
	if u.PassportNumber != "" {
		query += "[passport_number] = ?, "
		params = append(params, u.PassportNumber)
		hasUpdates = true
	}
	if u.PassportIssuedBy != "" {
		query += "[passport_issued_by] = ?, "
		params = append(params, u.PassportIssuedBy)
		hasUpdates = true
	}
	if u.PassportIssueDate != "" {
		query += "[passport_issue_date] = ?, "
		params = append(params, u.PassportIssueDate)
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
