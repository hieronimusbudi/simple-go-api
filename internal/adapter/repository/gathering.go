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
	gatheringAdapterRepository struct {
		db *sqlx.DB
	}

	GatheringAdapterRepositoryArgs struct {
		DB *sqlx.DB
	}
)

func NewGatheringRepository(args GatheringAdapterRepositoryArgs) repository.IGathering {
	return &gatheringAdapterRepository{
		db: args.DB,
	}
}

func (r *gatheringAdapterRepository) Create(ctx context.Context, gathering domain.Gathering) (id int64, err error) {
	query := `INSERT INTO gatherings (
		creator
		, type
		, scheduled_at
		, name
		, location
		, created_at
	) VALUES (?, ?, ?, ?, ?, NOW())`
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return
	}
	insertResult, err := tx.ExecContext(
		ctx,
		query,
		gathering.Creator.ID,
		gathering.Type,
		gathering.ScheduledAt,
		gathering.Name,
		gathering.Location,
	)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}
	id, err = insertResult.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Println(err)
	}
	if len(gathering.Attendees) > 0 {
		for _, member := range gathering.Attendees {
			err = createAttendee(ctx, tx, member.ID, id)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
	tx.Commit()
	return
}

func (r *gatheringAdapterRepository) Get(ctx context.Context, args domain.GatheringArgs) (gatherings []domain.Gathering, err error) {
	gatherings = []domain.Gathering{}
	conditions := []string{}
	query := `
		SELECT
			id
			, creator
			, type
			, scheduled_at
			, name
			, location
			, created_at
			, COALESCE(discarded_at, '') AS discarded_at
		FROM gatherings
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
	err = r.db.SelectContext(ctx, &gatherings, query)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return
	}

	attendeesQuery := `SELECT member_id, gathering_id FROM attendees`
	conditions = []string{}
	if len(args.IDs) > 0 {
		conditions = append(conditions, fmt.Sprintf(`gathering_id IN (%s)`, helpers.IntSliceToString(args.IDs)))
	}
	if len(conditions) > 0 {
		attendeesQuery += fmt.Sprintf(` WHERE %s`, strings.Join(conditions, " AND "))
	}
	rows, err := r.db.Query(
		attendeesQuery,
	)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return
	}
	defer rows.Close()
	mapAttendeesByGatheringID := map[int64][]int64{}
	for rows.Next() {
		var gID, aID int64
		if err = rows.Scan(
			&aID,
			&gID,
		); err != nil {
			log.Println(err)
			return
		}
		attendees, ok := mapAttendeesByGatheringID[gID]
		if !ok {
			attendees = []int64{}
		}
		attendees = append(attendees, aID)
		mapAttendeesByGatheringID[gID] = attendees
	}

	for i, g := range gatherings {
		g.Creator.ID = g.CreatorID
		attendees, ok := mapAttendeesByGatheringID[g.ID]
		if ok {
			g.Attendees = []domain.Member{}
			for _, a := range attendees {
				g.Attendees = append(g.Attendees, domain.Member{ID: a})
			}
		}
		gatherings[i] = g
	}
	return
}

func (r *gatheringAdapterRepository) Update(ctx context.Context, gathering domain.Gathering) (err error) {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return
	}

	query := `UPDATE gatherings SET
		type = ?
		, scheduled_at = ?
		, name = ?
		, location = ?
		, updated_at = NOW()
		WHERE id = ?`
	_, err = tx.ExecContext(
		ctx,
		query,
		gathering.Type,
		gathering.ScheduledAt,
		gathering.Name,
		gathering.Location,
		gathering.ID,
	)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}
	tx.Commit()
	return
}

func (r *gatheringAdapterRepository) Delete(ctx context.Context, args domain.GatheringArgs) (err error) {
	query := `UPDATE gatherings SET
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
