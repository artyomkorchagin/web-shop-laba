package psqlCategory

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
	"fmt"
)

func (r *Repository) AddCategory(ctx context.Context, ccr *types.CreateCategoryRequest) error {
	// Query the database to fetch all categories
	query := `
        INSERT INTO categories (
            name
        ) VALUES (
            $1
        )
    `

	_, err := r.db.QueryContext(ctx, query, ccr.Name)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
