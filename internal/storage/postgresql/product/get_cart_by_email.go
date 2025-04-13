package psqlProduct

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (r *Repository) GetCartByUserEmail(ctx context.Context, email string) ([]types.Product, error) {
	var products []types.Product
	var product types.Product
	query := `
        SELECT 
            p.product_id,
            p.name,
            p.description,
            p.price,
            p.image_url,
            ci.quantity
        FROM users u
        JOIN carts c ON u.user_id = c.user_id
        JOIN cart_items ci ON c.cart_id = ci.cart_id
        JOIN products p ON ci.product_id = p.product_id
        WHERE u.email = $1
    `

	rows, err := r.db.QueryContext(ctx, query, email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.ImageURL, &product.StockQuantity); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
