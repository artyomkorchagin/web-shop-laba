package psqlOrder

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
	"fmt"
)

func (r *Repository) GetOrderHistory(ctx context.Context, email string) ([]types.Order, error) {
	// Query to call the PostgreSQL function `get_orders`
	query := `SELECT * FROM get_orders($1)`

	// Execute the query
	rows, err := r.db.QueryContext(ctx, query, email)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// Map to store orders grouped by order_id
	orderMap := make(map[int]*types.Order)

	// Iterate over the rows returned by the query
	for rows.Next() {
		var (
			orderID            int
			imageURL           string
			status             string
			orderDate          string
			productName        string
			productDescription string
			categoryName       string
			quantity           int
			totalPrice         float64
			shippingAddress    string
			trackingNumber     string
		)

		// Scan the row into variables
		err := rows.Scan(
			&orderID,
			&productName,
			&productDescription,
			&categoryName,
			&quantity,
			&totalPrice,
			&shippingAddress,
			&trackingNumber,
			&orderDate,
			&status,
			&imageURL,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		order, exists := orderMap[orderID]
		if !exists {
			order = &types.Order{
				OrderID:         orderID,
				ShippingAddress: shippingAddress,
				TrackingNumber:  trackingNumber,
				OrderDate:       orderDate,
				Status:          status,
				Products:        []types.Product{},
			}
			orderMap[orderID] = order
		}

		order.Products = append(order.Products, types.Product{
			Name:           productName,
			StockQuantity:  quantity,
			Description:    productDescription,
			CategoryString: categoryName,
			ImageURL:       imageURL,
			Sum:            int(totalPrice),
		})
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	var orders []types.Order
	for _, order := range orderMap {
		orders = append(orders, *order)
	}

	return orders, nil
}
