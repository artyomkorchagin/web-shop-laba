package psqlCart

import (
	"context"
)

func (r *Repository) AddToCart(ctx context.Context, userID uint, productID uint) error {
	_, err := r.db.Exec("CALL add_to_cart($1, $2)", userID, productID)
	if err != nil {
		return err
	}
	return nil
}
