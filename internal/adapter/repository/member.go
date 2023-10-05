package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/hieronimusbudi/simple-go-api/internal/domain"
	"github.com/hieronimusbudi/simple-go-api/internal/domain/repository"
	"github.com/hieronimusbudi/simple-go-api/internal/helpers"
	"github.com/jmoiron/sqlx"
)

type (
	memberAdapterRepository struct {
		db *sqlx.DB
	}

	MemberAdapterRepositoryArgs struct {
		DB *sqlx.DB
	}
)

func NewMemberRepository(args MemberAdapterRepositoryArgs) repository.IMember {
	return &memberAdapterRepository{
		db: args.DB,
	}
}

func (r *memberAdapterRepository) Create(ctx context.Context, member domain.Member) (id int64, err error) {
	query := `INSERT INTO members (
		first_name
		, last_name
		, email
		, created_at
	) VALUES (?, ?, ?, NOW())`
	insertResult, err := r.db.ExecContext(
		ctx,
		query,
		member.FirstName,
		member.LastName,
		member.Email,
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

func (r *memberAdapterRepository) Get(ctx context.Context, args domain.MemberArgs) (members []domain.Member, err error) {
	members = []domain.Member{}
	conditions := []string{}
	query := `
		SELECT
			id
			, first_name
			, last_name
			, email
			, created_at
			, COALESCE(discarded_at, '') AS discarded_at
		FROM members
	`
	if !args.IsIncludeDiscard {
		conditions = append(conditions, `discarded_at IS NULL`)
	}
	if len(args.IDs) > 0 {
		conditions = append(conditions, fmt.Sprintf(`id IN (%s)`, helpers.IntSliceToString(args.IDs)))
	}
	if len(conditions) > 0 {
		query += fmt.Sprintf(` WHERE %s`, strings.Join(conditions, " AND "))
	}
	err = r.db.SelectContext(ctx, &members, query)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
	}
	return
}

func (r *memberAdapterRepository) Update(ctx context.Context, member domain.Member) (err error) {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return
	}
	query := `UPDATE members SET
		first_name = ?
		, last_name = ?
		, email = ?
		, updated_at = NOW()
		WHERE id = ?`
	_, err = tx.ExecContext(
		ctx,
		query,
		member.FirstName,
		member.LastName,
		member.Email,
		member.ID,
	)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}
	tx.Commit()
	return
}

func (r *memberAdapterRepository) Delete(ctx context.Context, args domain.MemberArgs) (err error) {
	query := `UPDATE members SET
		discarded_at = NOW()
		WHERE id = ?`
	_, err = r.db.ExecContext(
		ctx,
		query,
		args.ID,
	)
	if err != nil {
		log.Println(err)
	}
	return
}
