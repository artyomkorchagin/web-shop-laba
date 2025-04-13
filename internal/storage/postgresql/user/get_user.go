package psqlUser

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

func (r *Repository) GetUser(ctx context.Context, email string) (*types.User, error) {
	// Query the database to fetch the user by email
	query := `
        SELECT 
            user_id,
            username,
            date_of_birth,
            password_hash,
            role
        FROM users
        WHERE email = $1
    `
	var user types.User
	user.Email = email
	row := r.db.QueryRowContext(ctx, query, email)

	// Scan the result into the User struct
	err := row.Scan(
		&user.UserID,
		&user.Username,
		&user.DateOfBirth,
		&user.PasswordHash,
		&user.Role,
	)

	// Handle potential errors
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, errors.Wrap(err, "failed to execute query")
	}

	// Return the user object
	return &user, nil
}
