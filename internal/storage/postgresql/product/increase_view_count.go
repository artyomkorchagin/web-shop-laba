package psqlProduct

import "context"

func (r *Repository) IncreaseViewCount(ctx context.Context, productID int) error {
	query := `CALL increase_view_count($1)`
	_, err := r.db.ExecContext(ctx, query, productID)
	return err
}
