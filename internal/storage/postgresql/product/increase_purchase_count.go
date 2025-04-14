package psqlProduct

import "context"

func (r *Repository) IncreasePurchaseCount(ctx context.Context, productID int) error {
	query := `CALL increase_purchase_count($1)`
	_, err := r.db.ExecContext(ctx, query, productID)
	return err
}
