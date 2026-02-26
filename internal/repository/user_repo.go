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

func (r *UserRepo) CreateUserWithAuth(ctx context.Context, name, email, passwordHash, role string, dob time.Time) error {
	return r.q.CreateUserWithAuth(ctx, sqlc.CreateUserWithAuthParams{
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
		Dob:          dob,
	})
}

func (r *UserRepo) Get(ctx context.Context, id int32) (sqlc.User, error) {
	return r.q.GetUser(ctx, id)
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*sqlc.GetUserByEmailRow, error) {
	row, err := r.q.GetUserByEmail(ctx, email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &row, err
}

func (r *UserRepo) GetUserByID(ctx context.Context, id int32) (*sqlc.GetUserByIDRow, error) {
	row, err := r.q.GetUserByID(ctx, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &row, err
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

func (r *UserRepo) UpdateWithEmail(ctx context.Context, id int32, name, email string, dob time.Time) error {
	return r.q.UpdateUserWithEmail(ctx, sqlc.UpdateUserWithEmailParams{
		Name:  name,
		Email: email,
		Dob:   dob,
		ID:    id,
	})
}

func (r *UserRepo) Delete(ctx context.Context, id int32) error {
	return r.q.DeleteUser(ctx, id)
}
