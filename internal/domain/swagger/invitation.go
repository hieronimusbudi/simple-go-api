package swaggermodel

import (
	"github.com/hieronimusbudi/simple-go-api/internal/domain/valueobject"
)

type (
	Invitation struct {
		ID int64 `json:"id"`
		// Invitation status
		// * 0 -> Created
		// * 1 -> Accepted
		// * 2 -> Rejected
		// * 3 -> Cancelled
		Status    valueobject.InvitationStatus `json:"status" validate:"optional"`
		Member    MemberPayload                `json:"member" validate:"required"`
		Gathering GatheringPayload             `json:"gathering" validate:"required"`
		CreatedAt string                       `json:"created_at"`
	}
)
