package domain

import (
	"errors"
	"net/mail"
)

type (
	Member struct {
		ID          int64  `json:"id" db:"id"`
		FirstName   string `json:"first_name" db:"first_name"`
		LastName    string `json:"last_name" db:"last_name"`
		Email       string `json:"email" db:"email"`
		CreatedAt   string `json:"created_at" db:"created_at"`
		DiscardedAt string `json:"discarded_at,omitempty" db:"discarded_at"`
	}

	MemberArgs struct {
		IDs              []int64
		ID               int64
		IsIncludeDiscard bool
	}
)

func (d *Member) Validate() (err error) {
	if d.Email == "" {
		return errors.New("email is required")
	} else if _, err := mail.ParseAddress(d.Email); err != nil {
		return errors.New("invalid email format")
	}
	if d.FirstName == "" {
		return errors.New("first name is required")
	}
	return
}
