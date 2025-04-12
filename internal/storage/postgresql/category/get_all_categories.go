package psqlCategory

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
	"fmt"
)

func (r *Repository) GetAllCategories(ctx context.Context) ([]types.Category, error) {
	query := `
        SELECT 
            name
        FROM categories
        ORDER BY name ASC
    `

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var categories []types.Category
	for rows.Next() {
		var c types.Category
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		categories = append(categories, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return categories, nil
}
