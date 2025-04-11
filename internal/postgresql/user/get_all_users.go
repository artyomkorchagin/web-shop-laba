package mssqlUser

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
	"fmt"
)

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]types.User, error) {
	// Define the SQL query
	query := `
        SELECT 
            user_id,
			Email
        FROM Users
		WHERE role = 'default'
		`

	// Create a slice to store the users
	var users []types.User

	// Execute the query
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// Iterate over the rows and map them to User structs
	for rows.Next() {
		var user types.User
		if err := rows.Scan(&user.UserID, &user.Email); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		users = append(users, user)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	// If no users are found, return an empty slice
	if len(users) == 0 {
		return []types.User{}, nil
	}

	return users, nil
}
