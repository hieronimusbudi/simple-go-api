package domain

import (
	"errors"

	"github.com/hieronimusbudi/simple-go-api/internal/domain/valueobject"
)

type (
	Invitation struct {
		ID          int64                        `json:"id" db:"id"`
		MemberID    int64                        `json:"-" db:"member_id"`
		GatheringID int64                        `json:"-" db:"gathering_id"`
		Status      valueobject.InvitationStatus `json:"status" db:"status"`
		Member      Member                       `json:"member"`
		Gathering   Gathering                    `json:"gathering"`
		CreatedAt   string                       `json:"created_at" db:"created_at"`
	}

	InvitationArgs struct {
		IDs         []int64
		ID          int64
		MemberID    int64
		GatheringID int64
		Status      valueobject.InvitationStatus
	}
)

func (d *Invitation) Validate() (err error) {
	if d.Member.ID <= 0 {
		return errors.New("member is required")
	}
	if d.Gathering.ID <= 0 {
		return errors.New("gathering is required")
	}
	return
}
