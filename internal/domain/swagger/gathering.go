package swaggermodel

import (
	"github.com/hieronimusbudi/simple-go-api/internal/domain/valueobject"
)

type (
	Gathering struct {
		ID      int64         `json:"id" db:"id"`
		Creator MemberPayload `json:"creator" validate:"required"`
		// Gathering type
		// * 0 -> Private
		// * 1 -> Public
		Type valueobject.GatheringType `json:"type" db:"type" validate:"required"  example:"1"`
		// Date using (YYYY-MM-DD MM:SS) format
		ScheduledAt string          `json:"scheduled_at" db:"scheduled_at" validate:"required" example:"2023-10-06 04:53"`
		Name        string          `json:"name" db:"name" validate:"required" example:"Gathering Name"`
		Location    string          `json:"location" db:"location" validate:"required" example:"gathering street"`
		Attendees   []MemberPayload `json:"attendees" validate:"optional"`
	}

	UpdateGathering struct {
		// Gathering type
		// * 0 -> Private
		// * 1 -> Public
		Type valueobject.GatheringType `json:"type" db:"type" validate:"required"  example:"1"`
		// Date using (YYYY-MM-DD MM:SS) format
		ScheduledAt string `json:"scheduled_at" db:"scheduled_at" validate:"required" example:"2023-10-06 04:53"`
		Name        string `json:"name" db:"name" validate:"required" example:"Gathering Name"`
		Location    string `json:"location" db:"location" validate:"required" example:"gathering street"`
	}

	GatheringPayload struct {
		ID int64 `json:"id" validate:"required" example:"1"`
	}
)
