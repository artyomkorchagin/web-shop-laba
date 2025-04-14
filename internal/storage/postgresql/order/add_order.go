package psqlOrder

import (
	"context"
	"fmt"
)

func (r *Repository) AddOrder(ctx context.Context, email, address string) error {
	query := `CALL add_order($1, $2)`

	_, err := r.db.ExecContext(ctx, query, email, address)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
