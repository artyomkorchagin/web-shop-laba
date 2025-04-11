package psqlUser

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
			email
        FROM users
		WHERE role = 'customer'
		`

	var users []types.User

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user types.User
		if err := rows.Scan(&user.UserID, &user.Email); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	if len(users) == 0 {
		return []types.User{}, nil
	}

	return users, nil
}
