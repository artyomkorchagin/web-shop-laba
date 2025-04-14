package psqlProduct

import (
	"context"
	"fmt"
)

func (r *Repository) AddToCart(ctx context.Context, email string, productID, amount int) error {
	_, err := r.db.ExecContext(ctx, "SELECT add_to_cart($1, $2, $3)", email, productID, amount)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
