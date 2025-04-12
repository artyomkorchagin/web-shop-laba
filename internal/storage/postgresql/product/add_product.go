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
            image_url
        ) VALUES (
            $1, $2, $3, $4
        )
    `

	_, err := r.db.ExecContext(ctx, query,
		p.Name,
		p.Description,
		p.Price,
	)
	if err != nil {
		return fmt.Errorf("failed to insert product: %w", err)
	}

	return nil
}
