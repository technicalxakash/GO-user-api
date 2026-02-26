package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/akash/go-user-api/db/sqlc"
)

type UserRepo struct {
	q *sqlc.Queries
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		q: sqlc.New(db),
	}
}

func (r *UserRepo) Create(ctx context.Context, name string, dob time.Time) error {
	return r.q.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}

func (r *UserRepo) Get(ctx context.Context, id int32) (sqlc.User, error) {
	return r.q.GetUser(ctx, id)
}

func (r *UserRepo) List(ctx context.Context) ([]sqlc.User, error) {
	return r.q.ListUsers(ctx)
}

func (r *UserRepo) Update(ctx context.Context, id int32, name string, dob time.Time) error {
	return r.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		Name: name,
		Dob:  dob,
		ID:   id,
	})
}

func (r *UserRepo) Delete(ctx context.Context, id int32) error {
	return r.q.DeleteUser(ctx, id)
}