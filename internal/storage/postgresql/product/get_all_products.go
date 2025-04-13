package psqlProduct

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
	"fmt"
)

func (r *Repository) GetAllProducts(ctx context.Context) ([]types.Product, error) {
	query := `
        SELECT 
            product_id, 
            name, 
            description, 
            price,
			category_id,
			stock_quantity,
			image_url
        FROM products
        ORDER BY created_at DESC
    `

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var products []types.Product
	for rows.Next() {
		var p types.Product
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Category,
			&p.StockQuantity,
			&p.ImageURL,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return products, nil
}
