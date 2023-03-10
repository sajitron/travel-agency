// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateSession(ctx context.Context, arg CreateSessionParams) (Sessions, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
	GetSession(ctx context.Context, id uuid.UUID) (Sessions, error)
	GetUser(ctx context.Context, email string) (Users, error)
	GetUserById(ctx context.Context, id int64) (Users, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error)
}

var _ Querier = (*Queries)(nil)
