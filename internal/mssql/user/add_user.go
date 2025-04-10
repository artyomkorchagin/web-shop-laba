package mssqlUser

import (
	"context"
	"socialsecurity/internal/types"
)

func (r *mssqlUserRepository) AddUser(ctx context.Context, user *types.User) error {
	query := `
        INSERT INTO [dbo].[Users] (
            [first_name],
            [second_name],
            [middle_name],
            [date_of_birth],
            [gender],
            [address],
            [phone_number],
            [email],
            [password_hash],
            [snils],
            [passport_series],
            [passport_number],
            [passport_issued_by],
            [passport_issue_date]
        ) VALUES (
            ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
        )
    `

	_, err := r.db.ExecContext(ctx, query,
		user.FirstName,
		user.SecondName,
		user.MiddleName,
		user.DateOfBirth,
		user.Gender,
		user.Address,
		user.PhoneNumber,
		user.Email,
		user.PasswordHash,
		user.SNILS,
		user.PassportSeries,
		user.PassportNumber,
		user.PassportIssuedBy,
		user.PassportIssueDate,
	)

	return err
}
