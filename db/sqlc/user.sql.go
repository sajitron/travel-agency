// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    first_name,
    last_name,
    email,
    password
) VALUES (
    $1, $2, $3, $4
) RETURNING id, first_name, last_name, email, password, password_changed_at, created_at, updated_at
`

type CreateUserParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
	)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, first_name, last_name, email, password, password_changed_at, created_at, updated_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (Users, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, first_name, last_name, email, password, password_changed_at, created_at, updated_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id int64) (Users, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
  password = COALESCE($1, password),
  password_changed_at = COALESCE($2, password_changed_at),
  first_name = COALESCE($3, first_name),
  last_name = COALESCE($4, last_name),
  email = COALESCE($5, email)
WHERE
  id = $6
RETURNING id, first_name, last_name, email, password, password_changed_at, created_at, updated_at
`

type UpdateUserParams struct {
	Password          sql.NullString `json:"password"`
	PasswordChangedAt sql.NullTime   `json:"password_changed_at"`
	FirstName         sql.NullString `json:"first_name"`
	LastName          sql.NullString `json:"last_name"`
	Email             sql.NullString `json:"email"`
	ID                int64          `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Password,
		arg.PasswordChangedAt,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.ID,
	)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
