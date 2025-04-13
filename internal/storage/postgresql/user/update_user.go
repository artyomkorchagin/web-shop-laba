package psqlUser

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func (r *Repository) UpdateUser(ctx context.Context, u *types.User) error {
	if u.Username == "" && u.Email == "" {
		return errors.New("at least one field (username or email) must be provided for update")
	}

	query := "UPDATE users SET "
	args := []interface{}{}
	setClauses := []string{}

	if u.Username != "" {
		setClauses = append(setClauses, "username = $1")
		args = append(args, u.Username)
	}
	if u.Email != "" {
		setClauses = append(setClauses, fmt.Sprintf("email = $%d", len(args)+1))
		args = append(args, u.Email)
	}

	query += strings.Join(setClauses, ", ") + " WHERE email = $" + fmt.Sprint(len(args)+1)
	args = append(args, u.Email)

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("no user found with the given ID")
	}

	return nil
}
