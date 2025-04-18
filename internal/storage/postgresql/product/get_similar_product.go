package psqlProduct

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
	"fmt"
)

func (r *Repository) GetSimilarProduct(ctx context.Context, userInput string) (types.Product, error) {

	var product types.Product

	query := `
        SELECT 
            product_id,
            name
        FROM products
        WHERE name ILIKE $1 AND stock_quantity > 0
        LIMIT 1;
    `

	searchTerm := "%" + userInput + "%"

	if err := r.db.
		QueryRowContext(ctx, query, searchTerm).
		Scan(&product.ID, &product.Name); err != nil {
		return product, fmt.Errorf("Такого товара нет")
	}
	return product, nil
}
