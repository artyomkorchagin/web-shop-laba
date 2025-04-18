package psqlProduct

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (r *Repository) GetProductById(ctx context.Context, id int) (types.Product, error) {
	var product types.Product
	query := `
		SELECT
			p.name,
			p.price,
			p.description,
			p.stock_quantity,
			c.name,
			p.created_at,
			p.image_url
		FROM products p
		JOIN categories c ON p.category_id = c.category_id
		WHERE p.product_id = $1
	`
	if err := r.db.
		QueryRowContext(ctx, query, id).
		Scan(&product.Name, &product.Price, &product.Description, &product.StockQuantity, &product.CategoryString, &product.CreatedAt, &product.ImageURL); err != nil {
		return product, err
	}
	product.ID = id
	return product, nil
}
