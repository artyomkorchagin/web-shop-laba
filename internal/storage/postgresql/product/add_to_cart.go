package psqlProduct

import (
	"context"
)

func (r *Repository) AddToCart(ctx context.Context, email string, productID, amount int) error {
	_, err := r.db.Exec("CALL add_to_cart($1, $2)", email, productID)
	if err != nil {
		return err
	}
	return nil
}
