package mssqlApplication

import (
	"context"
	"fmt"
	utils "socialsecurity/internal/mssql"
	"socialsecurity/internal/types"

	"github.com/google/uuid"
)

func (r *mssqlApplicationRepository) GetServices(ctx context.Context) ([]types.Service, error) {
	// Define the SQL query
	query := `
        SELECT 
            service_id, 
            service_name, 
            description, 
            provider 
        FROM Service`

	// Create a slice to store the services
	var services []types.Service

	// Execute the query
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// Iterate over the rows and map them to Service structs
	for rows.Next() {
		var service types.Service
		if err := rows.Scan(&service.ServiceID, &service.Name, &service.Description, &service.Provider); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		normUUID := utils.ReverseUUID(service.ServiceID.String())

		service.ServiceID, err = uuid.Parse(normUUID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse service ID: %w", err)
		}

		services = append(services, service)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	// If no services are found, return an empty slice
	if len(services) == 0 {
		return []types.Service{}, nil
	}

	return services, nil
}
