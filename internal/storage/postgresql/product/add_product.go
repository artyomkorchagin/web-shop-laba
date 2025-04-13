package psqlProduct

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
	"fmt"
)

func (r *Repository) AddProduct(ctx context.Context, p *types.Product) error {
	query := `
        INSERT INTO products (
            name, 
            description, 
            price, 
			stock_quantity,
			category_id,
            image_url
        ) VALUES (
            $1, $2, $3, $4, $5, $6
        )
    `

	_, err := r.db.ExecContext(ctx, query,
		p.Name,
		p.Description,
		p.Price,
		p.StockQuantity,
		p.Category,
		p.ImageURL,
	)
	if err != nil {
		fmt.Print(err)
		return fmt.Errorf("failed to insert product: %w", err)
	}

	return nil
}
