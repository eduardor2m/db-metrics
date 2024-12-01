package postgres

import (
	"context"
	"fmt"

	"github.com/eduardor2m/db-metrics/src/internal/models"
	"github.com/eduardor2m/db-metrics/src/internal/repositories/postgres/bridge"
	"github.com/sqlc-dev/pqtype"
)

type UserPostgresRepository struct {
	connectorManager
}

func (instance UserPostgresRepository) Create(u models.User) error {
	conn, err := instance.getConnection()

	if err != nil {
		return fmt.Errorf("error getting connection: %v", err)
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	err = queries.CreateUser(ctx, bridge.CreateUserParams{
		ID:          u.ID(),
		Name:        u.Name(),
		Email:       u.Email(),
		Preferences: pqtype.NullRawMessage{Valid: true, RawMessage: []byte("[]")},
	})

	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	return nil
}

func (instance UserPostgresRepository) List() ([]models.User, error) {
	conn, err := instance.getConnection()

	if err != nil {
		return nil, fmt.Errorf("error getting connection: %v", err)
	}

	defer instance.closeConnection(conn)

	ctx := context.Background()

	queries := bridge.New(conn)

	users, err := queries.ListUsers(ctx)

	if err != nil {
		return nil, fmt.Errorf("error listing users: %v", err)
	}

	var result []models.User

	for _, user := range users {
		var preferences []string = nil

		result = append(result, *models.NewUser(user.ID, user.Name, user.Email, preferences))
	}

	return result, nil
}
