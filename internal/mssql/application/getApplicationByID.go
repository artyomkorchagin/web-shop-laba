package mssqlApplication

import (
	"context"
	"database/sql"
	"fmt"
	"socialsecurity/internal/types"

	"github.com/google/uuid"
)

func (r *mssqlApplicationRepository) GetApplicationByID(ctx context.Context, id uuid.UUID) (*types.Application, error) {
	// Define the SQL query
	query := `
        SELECT 
            user_id, 
            benefit_id, 
            service_id, 
            application_date, 
            status, 
            approval_date, 
            rejection_reason, 
            description
        FROM applications
        WHERE application_id = ?
    `

	// Execute the query
	row := r.db.QueryRowContext(ctx, query, id)

	// Create an Application struct to store the result
	var app types.Application

	// Handle nullable fields using sql.NullString
	var approvalDate sql.NullString
	var rejectionReason sql.NullString
	var description sql.NullString

	// Scan the row into the struct fields
	err := row.Scan(
		&app.UserID,
		&app.BenefitID,
		&app.ServiceID,
		&app.Date,
		&app.Status,
		&approvalDate,
		&rejectionReason,
		&description,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("application not found")
		}
		return nil, fmt.Errorf("failed to retrieve application: %w", err)
	}

	// Assign nullable fields if they are valid

	if approvalDate.Valid {
		app.ApprovalDate = approvalDate.String
	}

	if rejectionReason.Valid {
		app.RejectionReason = rejectionReason.String
	}

	if description.Valid {
		app.Description = description.String
	}
	app.ID = id
	fmt.Println(app)
	return &app, nil
}
