package history

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// OneEntry is a helper that creates a new history position with one entry.
//
// Use NewPosition to create a position with more then one entry.
func OneEntry(ctx context.Context, tx pgx.Tx, userID int, modelID string, meetingID int, entryList ...string) error {
	pos, err := NewPosition(ctx, tx, userID)
	if err != nil {
		return fmt.Errorf("creating new position: %w", err)
	}

	if err := pos.NewEntry(ctx, modelID, meetingID, entryList...); err != nil {
		return fmt.Errorf("write entry: %w", err)
	}

	return nil
}

// Position represents a position in the history for reading.
type Position struct {
	ID int
	tx pgx.Tx
}

// NewPosition creates a new position in the history.
func NewPosition(ctx context.Context, tx pgx.Tx, userID int) (Position, error) {
	sql := `INSERT INTO history_position_t (timestamp, original_user_id, user_id)
            VALUES (NOW(), $1, $1)
            RETURNING id`

	pos := Position{
		tx: tx,
	}
	if err := tx.QueryRow(ctx, sql, userID).Scan(&pos.ID); err != nil {
		return Position{}, fmt.Errorf("insert new position: %w", err)
	}

	return pos, nil
}

// NewEntry adds a new entry to the history at this position.
func (p *Position) NewEntry(ctx context.Context, modelID string, meetingID int, entryList ...string) error {
	sql := `INSERT INTO history_entry_t (entries, original_model_id, model_id, position_id, meeting_id)
            VALUES ($1, $2, $2, $3, $4)`

	_, err := p.tx.Exec(ctx, sql, entryList, modelID, p.ID, meetingID)
	if err != nil {
		return fmt.Errorf("insert new entry: %w", err)
	}

	return nil
}
