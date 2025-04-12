package psqlCategory

import (
	"context"
	"fmt"
)

func (r *Repository) AddCategory(ctx context.Context, name string) error {
	// Query the database to fetch all categories
	query := `
        INSERT INTO categories (
            name
        ) VALUES (
            $1
        )
    `

	_, err := r.db.QueryContext(ctx, query, name)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
