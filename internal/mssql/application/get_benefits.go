package mssqlApplication

import (
	"context"
	"fmt"
	utils "socialsecurity/internal/mssql"
	"socialsecurity/internal/types"

	"github.com/google/uuid"
)

func (r *mssqlApplicationRepository) GetBenefits(ctx context.Context) ([]types.Benefit, error) {
	// Define the SQL query
	query := `
        SELECT 
            benefit_id, 
            benefit_name, 
            description, 
            amount,
			frequency
        FROM Benefits`

	// Create a slice to store the benefits
	var benefits []types.Benefit

	// Execute the query
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// Iterate over the rows and map them to Benefit structs
	for rows.Next() {
		var benefit types.Benefit
		if err := rows.Scan(&benefit.BenefitID, &benefit.Name, &benefit.Description, &benefit.Amount, &benefit.Frequency); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		normUUID := utils.ReverseUUID(benefit.BenefitID.String())

		benefit.BenefitID, err = uuid.Parse(normUUID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse benefit ID: %w", err)
		}

		benefits = append(benefits, benefit)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	// If no benefits are found, return an empty slice
	if len(benefits) == 0 {
		return []types.Benefit{}, nil
	}

	return benefits, nil
}
