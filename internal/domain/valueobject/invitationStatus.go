package valueobject

type InvitationStatus int

const (
	INVITATION_CREATED  InvitationStatus = 0
	INVITATION_ACCEPT   InvitationStatus = 1
	INVITATION_REJECT   InvitationStatus = 2
	INVITATION_CANCELED InvitationStatus = 3
)
