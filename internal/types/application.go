package types

import "github.com/google/uuid"

type Application struct {
	ID              uuid.UUID  `json:"application_id" binding:"required,uuid"`
	UserID          uuid.UUID  `json:"user_id" binding:"required,uuid"`
	BenefitID       *uuid.UUID `json:"benefit_id" binding:"omitempty,uuid"`
	ServiceID       *uuid.UUID `json:"service_id" binding:"omitempty,uuid"`
	Date            string     `json:"application_date"`
	Status          string     `json:"status"`
	ApprovalDate    string     `json:"approval_date"`
	RejectionReason string     `json:"rejection_reason"`
	Description     string     `json:"description"`
}

type CreateApplicationRequest struct {
	UserID      uuid.UUID `json:"user_id" binding:"required,uuid"`
	BenefitID   uuid.UUID `json:"benefit_id" binding:"omitempty,uuid"`
	ServiceID   uuid.UUID `json:"service_id" binding:"omitempty,uuid"`
	Description string    `json:"description"`
}

type WorkApplicationRequest struct {
	ID              uuid.UUID `json:"application_id"`
	Status          string    `json:"status"`
	RejectionReason string    `json:"rejection_reason"`
}

type EditApplicatonRequest struct {
	Description string    `json:"description"`
	BenefitID   uuid.UUID `json:"benefit_id"`
	ServiceID   uuid.UUID `json:"service_id"`
}

const (
	StatusPending   = "pending"
	StatusApproved  = "approved"
	StatusRejected  = "rejected"
	StatusCancelled = "cancelled"
)
