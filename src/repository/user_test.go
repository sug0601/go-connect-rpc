// src/repository/user_test.go
package repository

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"example.com/src/infra"
	"example.com/model"
	"github.com/stretchr/testify/require"
)

func setupUserRepoTest(t *testing.T) (*BunUserRepository, context.Context, func(*model.User)) {
	db := infra.NewPostgresDB("postgres://user:pass@localhost:5433/connect?sslmode=disable")
	repo := NewBunUserRepository(db)
	ctx := context.Background()

	cleanup := func(user *model.User) {
		if user != nil {
			_, err := db.NewDelete().
				Model(user).
				WherePK().
				Exec(ctx)
			if err != nil {
				t.Logf("failed to cleanup user: %v", err)
			}
		}
		db.Close()
	}

	return repo, ctx, cleanup
}

func randomEmail(name string) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s_%d@example.com", name, rand.Intn(1000000))
}

func TestUserRepository_Insert(t *testing.T) {
	repo, ctx, cleanup := setupUserRepoTest(t)

	user, err := repo.Insert(ctx, "Alice", randomEmail("Alice"))
	require.NoError(t, err)
	require.NotEmpty(t, user.ID)

	defer cleanup(user)
}

func TestUserRepository_FindByID(t *testing.T) {
	repo, ctx, cleanup := setupUserRepoTest(t)

	user, err := repo.Insert(ctx, "Bob", randomEmail("Bob"))
	require.NoError(t, err)
	defer cleanup(user)

	found, err := repo.FindByID(ctx, user.ID)
	require.NoError(t, err)
	require.Equal(t, "Bob", found.Name)
	require.Equal(t, user.Email, found.Email)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	repo, ctx, cleanup := setupUserRepoTest(t)

	user, err := repo.Insert(ctx, "Charlie", randomEmail("Charlie"))
	require.NoError(t, err)
	defer cleanup(user)

	found, err := repo.FindByEmail(ctx, user.Email)
	require.NoError(t, err)
	require.Equal(t, user.ID, found.ID)
}