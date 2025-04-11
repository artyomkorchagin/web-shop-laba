package pgsqlUser

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
	"time"
)

func (r *UserRepository) AddUser(ctx context.Context, user *types.User) error {
	query := `
        INSERT INTO users (
            username,
            email,
            password_hash,
            role,
            date_of_birth,
            created_at,
            last_login
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7
        )
    `

	// Execute the query with parameters
	_, err := r.db.ExecContext(ctx, query,
		user.Username,     // $1
		user.Email,        // $2
		user.PasswordHash, // $3
		user.Role,         // $4
		user.DateOfBirth,  // $5
		time.Now(),        // $6: created_at (current timestamp)
		nil,               // $7: last_login (set to NULL initially)
	)

	return err
}
