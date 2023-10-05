package repository

import (
	"context"
	"database/sql"
	"log"
)

func createAttendee(ctx context.Context, tx *sql.Tx, memberID int64, gatheringID int64) (err error) {
	_, err = tx.ExecContext(ctx, `
	INSERT INTO attendees (
		member_id
		, gathering_id
	) VALUES (?, ?)`, memberID, gatheringID)
	if err != nil {
		tx.Rollback()
		log.Println(err)
	}
	return
}

func removeAttendee(ctx context.Context, tx *sql.Tx, memberID int64, gatheringID int64) (err error) {
	_, err = tx.ExecContext(ctx, `DELETE FROM attendees WHERE member_id = ? AND gathering_id = ?`, memberID, gatheringID)
	if err != nil {
		tx.Rollback()
		log.Println(err)
	}
	return
}
