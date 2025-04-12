package psqlCart

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (r *Repository) GetCartByUserID(ctx context.Context, userid uint) (*types.Cart, error) {
	var cart types.Cart
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM carts WHERE userid=$1", userid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err = rows.Scan(&cart.Id, &cart.UserId, &cart.Products); err != nil {
			return nil, err
		}
	}
	return &cart, nil
}
