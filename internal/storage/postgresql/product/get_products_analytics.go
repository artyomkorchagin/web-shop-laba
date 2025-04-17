package psqlProduct

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
	"database/sql"
	"fmt"
)

func (r *Repository) GetProductsAnalytics(ctx context.Context) ([]types.Product, error) {
	query := `SELECT * FROM get_products_stats()`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Failed to execute query:", err)
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var products []types.Product
	for rows.Next() {
		var p types.Product
		var lastViewAt, lastPurchaseAt sql.NullString // Use sql.NullString for nullable fields

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.ImageURL,
			&p.ViewCount,
			&p.PurchaseCount,
			&lastViewAt,     // Nullable field
			&lastPurchaseAt, // Nullable field
		)
		if err != nil {
			fmt.Println("Failed to scan row:", err)
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Convert sql.NullString to *string
		if lastViewAt.Valid {
			p.LastViewAt = lastViewAt.String
		} else {
			p.LastViewAt = ""
		}

		if lastPurchaseAt.Valid {
			p.LastPurchaseAt = lastPurchaseAt.String
		} else {
			p.LastPurchaseAt = ""
		}

		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during row iteration:", err)
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return products, nil
}
