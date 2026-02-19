package auth

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// UserLookup provides user ID lookup by keycloak_id
type UserLookup struct {
	pool *pgxpool.Pool
}

// NewUserLookup creates a UserLookup with the given connection pool
func NewUserLookup(pool *pgxpool.Pool) *UserLookup {
	return &UserLookup{pool: pool}
}

// GetUserIDByKeycloakID returns the OpenSlides user ID for the given keycloak_id
func (l *UserLookup) GetUserIDByKeycloakID(ctx context.Context, keycloakID string) (int, error) {
	var userID int
	err := l.pool.QueryRow(ctx,
		"SELECT id FROM user_t WHERE keycloak_id = $1 AND is_active = true",
		keycloakID,
	).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("user lookup by keycloak_id: %w", err)
	}
	return userID, nil
}
