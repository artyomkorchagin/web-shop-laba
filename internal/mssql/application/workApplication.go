package mssqlApplication

import (
	"context"
	"database/sql"
	"fmt"
	utils "socialsecurity/internal/mssql"
	"socialsecurity/internal/types"

	"github.com/google/uuid"
)

func (r *mssqlApplicationRepository) WorkApplication(ctx context.Context, input types.WorkApplicationRequest) error {
	// Validate the input
	if input.ID == uuid.Nil {
		return fmt.Errorf("invalid application ID")
	}
	if input.Status != "approved" && input.Status != "rejected" {
		return fmt.Errorf("invalid status: %s. Allowed values are 'approved' or 'rejected'", input.Status)
	}
	if input.Status == "rejected" && input.RejectionReason == "" {
		return fmt.Errorf("rejection reason is required when status is 'rejected'")
	}
	appID := utils.ReverseUUID(input.ID.String())
	// Define the SQL query
	query := `
        UPDATE Applications
        SET 
            status = ?,
            rejection_reason = ?
        WHERE application_id = ?
    `

	fmt.Println(input)
	result, err := r.db.ExecContext(ctx, query, input.Status, sql.NullString{String: input.RejectionReason, Valid: input.Status == "rejected"}, appID)
	if err != nil {
		return fmt.Errorf("failed to update application: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("application not found")
	}

	return nil
}
