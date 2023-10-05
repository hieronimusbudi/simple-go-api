package swaggermodel

type (
	Member struct {
		ID        int64  `json:"id" db:"id"`
		FirstName string `json:"first_name" db:"first_name" validate:"required" example:"John"`
		LastName  string `json:"last_name" db:"last_name" example:"Doe"`
		Email     string `json:"email" db:"email" validate:"required" example:"john@mail.com"`
	}

	MemberPayload struct {
		ID int64 `json:"id" validate:"required" example:"1"`
	}
)
