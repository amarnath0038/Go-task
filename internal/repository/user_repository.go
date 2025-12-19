package repository

import (
	"context"
	"time"

	"user-crud/db/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository struct {
	queries *sqlc.Queries
}

func NewUserRepository(q *sqlc.Queries) *UserRepository {
	return &UserRepository{
		queries: q,
	}
}

func (r *UserRepository) CreateUser(
	ctx context.Context,
	name string,
	dob time.Time,
) (sqlc.User, error) {

	pgDob := pgtype.Date{
		Time:  dob,
		Valid: true,
	}

	return r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  pgDob,
	})
}

func (r *UserRepository) GetUser(
	ctx context.Context,
	id int64,
) (sqlc.User, error) {
	return r.queries.GetUser(ctx, int32(id))
}

func (r *UserRepository) ListUsers(
	ctx context.Context,
	limit int32,
	offset int32,
) ([]sqlc.User, error) {
	return r.queries.ListUsers(ctx, sqlc.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
}

func (r *UserRepository) UpdateUser(
	ctx context.Context,
	id int64,
	name string,
	dob time.Time,
) (sqlc.User, error) {

	pgDob := pgtype.Date{
		Time:  dob,
		Valid: true,
	}

	return r.queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   int32(id),
		Name: name,
		Dob:  pgDob,
	})
}

func (r *UserRepository) DeleteUser(
	ctx context.Context,
	id int64,
) error {
	return r.queries.DeleteUser(ctx, int32(id))
}
