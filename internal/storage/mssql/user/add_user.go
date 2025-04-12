package mssqlUser

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
	"time"
)

func (r *Repository) AddUser(ctx context.Context, user *types.User) error {
	query := `
        INSERT INTO [dbo].[Users] (
            ,
        ) VALUES (
            ?, ?, ?, 
        )
    `

	_, err := r.db.ExecContext(ctx, query,
		user.Username,
		user.Email,
		user.PasswordHash,
		time.Now(),
	)

	return err
}
