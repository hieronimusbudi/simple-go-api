package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/hieronimusbudi/simple-go-api/internal/domain"
	"github.com/hieronimusbudi/simple-go-api/internal/domain/repository"
	"github.com/hieronimusbudi/simple-go-api/internal/domain/valueobject"
	"github.com/hieronimusbudi/simple-go-api/internal/helpers"
	"github.com/jmoiron/sqlx"
)

type (
	invitationAdapterRepository struct {
		db *sqlx.DB
	}

	InvitationAdapterRepositoryArgs struct {
		DB *sqlx.DB
	}
)

func NewInvitationRepository(args InvitationAdapterRepositoryArgs) repository.IInvitation {
	return &invitationAdapterRepository{
		db: args.DB,
	}
}

func (r *invitationAdapterRepository) Create(ctx context.Context, invitation domain.Invitation) (id int64, err error) {
	query := `INSERT INTO invitations (
		member_id
		, gathering_id
		, status
		, created_at
	) VALUES (?, ?, ?, NOW())`
	insertResult, err := r.db.ExecContext(
		ctx,
		query,
		invitation.Member.ID,
		invitation.Gathering.ID,
		invitation.Status,
	)
	if err != nil {
		log.Println(err)
		return
	}
	id, err = insertResult.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return
}

func (r *invitationAdapterRepository) Get(ctx context.Context, args domain.InvitationArgs) (invitations []domain.Invitation, err error) {
	invitations = []domain.Invitation{}
	conditions := []string{}
	query := `
		SELECT
			id
			, member_id
			, gathering_id
			, status
			, created_at
		FROM invitations
	`
	if len(args.IDs) > 0 {
		conditions = append(conditions, fmt.Sprintf(`id IN (%s)`, helpers.IntSliceToString(args.IDs)))
	}
	if len(conditions) > 0 {
		query += fmt.Sprintf(` WHERE %s`, strings.Join(conditions, " AND "))
	}
	err = r.db.SelectContext(ctx, &invitations, query)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
	}
	for i, inv := range invitations {
		inv.Member.ID = inv.MemberID
		inv.Gathering.ID = inv.GatheringID
		invitations[i] = inv
	}
	return
}

func (r *invitationAdapterRepository) UpdateStatus(ctx context.Context, args domain.InvitationArgs) (err error) {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return
	}
	query := `UPDATE invitations SET
		status = ?
		WHERE id = ?`
	_, err = tx.ExecContext(
		ctx,
		query,
		args.Status,
		args.ID,
	)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}
	if args.Status == valueobject.INVITATION_ACCEPT {
		err = createAttendee(ctx, tx, args.MemberID, args.GatheringID)
		if err != nil {
			tx.Rollback()
			if strings.Contains(err.Error(), "Error 1062: Duplicate entry") {
				err = errors.New("the member has accepted the invitation")
			}
			log.Println(err)
			return
		}
	} else if args.Status == valueobject.INVITATION_REJECT || args.Status == valueobject.INVITATION_CANCELED {
		err = removeAttendee(ctx, tx, args.MemberID, args.GatheringID)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return
		}
	}
	tx.Commit()
	return
}
