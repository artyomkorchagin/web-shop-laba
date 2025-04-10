package mssqlUser

import (
	"context"
	"database/sql"
	utils "socialsecurity/internal/mssql"
	"socialsecurity/internal/types"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (r *mssqlUserRepository) GetUser(ctx context.Context, email string) (*types.User, error) {
	// Query the database to fetch the user by email and hashed password
	query := `
        SELECT 
            *
        FROM [dbo].[Users]
        WHERE [email] = ?
    `

	// Execute the query
	var user types.User
	row := r.db.QueryRowContext(ctx, query, email)

	// Scan the result into the User struct
	err := row.Scan(
		&user.UserID,
		&user.FirstName,
		&user.SecondName,
		&user.MiddleName,
		&user.DateOfBirth,
		&user.Gender,
		&user.Address,
		&user.PhoneNumber,
		&user.Email,
		&user.PasswordHash,
		&user.SNILS,
		&user.PassportSeries,
		&user.PassportNumber,
		&user.PassportIssuedBy,
		&user.PassportIssueDate,
		&user.Role,
	)

	// Handle potential errors
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, errors.Wrap(err, "failed to execute query")
	}
	normUUID := utils.ReverseUUID(user.UserID.String())
	user.UserID, err = uuid.Parse(normUUID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse UUID")
	}
	// Return the user object
	return &user, nil
}
