package domain

import (
	"errors"
	"time"

	"github.com/hieronimusbudi/simple-go-api/internal/domain/valueobject"
)

type (
	Gathering struct {
		ID          int64                     `json:"id" db:"id"`
		CreatorID   int64                     `json:"-" db:"creator"`
		Creator     Member                    `json:"creator"`
		Type        valueobject.GatheringType `json:"type" db:"type"`
		ScheduledAt string                    `json:"scheduled_at" db:"scheduled_at"`
		Name        string                    `json:"name" db:"name"`
		Location    string                    `json:"location" db:"location"`
		Attendees   []Member                  `json:"attendees"`
		CreatedAt   string                    `json:"created_at" db:"created_at"`
		DiscardedAt string                    `json:"discarded_at,omitempty" db:"discarded_at"`
	}

	GatheringArgs struct {
		IDs              []int64
		MemberIDs        []int64
		ID               int64
		IsIncludeDiscard bool
	}
)

func (d *Gathering) Validate() (err error) {
	if d.Creator.ID <= 0 {
		return errors.New("creator is required")
	}
	if d.ScheduledAt == "" {
		return errors.New("scheduled at is required")
	} else {
		_, err = time.Parse("2006-01-02 15:04", d.ScheduledAt)
		if err != nil {
			return errors.New("invalid time format, please use (YYYY-MM-DD MM:SS) format")
		}
	}
	if d.Location == "" {
		return errors.New("location at is required")
	}
	if d.Name == "" {
		return errors.New("gathering name at is required")
	}
	if d.Type != valueobject.PRIVATE && d.Type != valueobject.PUBLIC {
		d.Type = valueobject.PRIVATE
	}
	if len(d.Attendees) > 0 {
		for _, a := range d.Attendees {
			if a.ID <= 0 {
				return errors.New("attendee is required")
			}
		}
	}
	return
}
