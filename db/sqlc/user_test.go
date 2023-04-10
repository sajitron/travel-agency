package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/sajitron/travel-agency/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) Users {
	password := util.RandomString(7)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)
	require.NotEqual(t, password, hashedPassword)

	arg := CreateUserParams{
		Email:     util.RandomEmail(),
		FirstName: util.RandomName(),
		LastName:  util.RandomName(),
		Password:  hashedPassword,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Password, user.Password)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	newUser := createRandomUser(t)
	retrievedUser, err := testQueries.GetUser(context.Background(), newUser.Email)

	require.NoError(t, err)
	require.NotEmpty(t, retrievedUser)

	require.Equal(t, newUser.Email, retrievedUser.Email)
	require.Equal(t, newUser.FirstName, retrievedUser.FirstName)
	require.Equal(t, newUser.LastName, retrievedUser.LastName)
	require.Equal(t, newUser.Password, retrievedUser.Password)
	require.WithinDuration(t, newUser.CreatedAt, retrievedUser.CreatedAt, time.Second)
	require.WithinDuration(t, newUser.UpdatedAt, retrievedUser.UpdatedAt, time.Second)
}

func TestGetUserByID(t *testing.T) {
	newUser := createRandomUser(t)
	retrievedUser, err := testQueries.GetUserById(context.Background(), newUser.ID)

	require.NoError(t, err)
	require.NotEmpty(t, retrievedUser)

	require.Equal(t, newUser.ID, retrievedUser.ID)
	require.Equal(t, newUser.Email, retrievedUser.Email)
	require.Equal(t, newUser.FirstName, retrievedUser.FirstName)
	require.Equal(t, newUser.LastName, retrievedUser.LastName)
	require.Equal(t, newUser.Password, retrievedUser.Password)
	require.WithinDuration(t, newUser.CreatedAt, retrievedUser.CreatedAt, time.Second)
	require.WithinDuration(t, newUser.UpdatedAt, retrievedUser.UpdatedAt, time.Second)
}

func TestUpdateUserFirstAndLastName(t *testing.T) {
	user := createRandomUser(t)

	newFirstName := util.RandomName()
	newLastName := util.RandomName()

	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		ID: user.ID,
		FirstName: sql.NullString{
			String: newFirstName,
			Valid:  true,
		},
		LastName: sql.NullString{
			String: newLastName,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, user, updatedUser)
	require.Equal(t, newFirstName, updatedUser.FirstName)
	require.Equal(t, newLastName, updatedUser.LastName)
	require.Equal(t, user.Email, updatedUser.Email)
	require.Equal(t, user.Password, updatedUser.Password)
}

func TestUpdatePassword(t *testing.T) {
	user := createRandomUser(t)

	newPassword := util.RandomString(8)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		ID: user.ID,
		Password: sql.NullString{
			String: newHashedPassword,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, user, updateUser)
	require.NotEqual(t, user.Password, updatedUser.Password)
	require.Equal(t, user.FirstName, updatedUser.FirstName)
	require.Equal(t, user.LastName, updatedUser.LastName)
	require.Equal(t, user.Email, updatedUser.Email)
	require.Equal(t, newHashedPassword, updatedUser.Password)
}

func TestUpdateEmail(t *testing.T) {
	user := createRandomUser(t)

	newEmail := util.RandomEmail()

	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		ID: user.ID,
		Email: sql.NullString{
			String: newEmail,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, user, updatedUser)
	require.NotEqual(t, user.Email, updatedUser.Email)
	require.Equal(t, user.FirstName, updatedUser.FirstName)
	require.Equal(t, user.LastName, updatedUser.LastName)
	require.Equal(t, user.Password, updatedUser.Password)
}

func TestUpdateAllFields(t *testing.T) {
	user := createRandomUser(t)

	newFirstName := util.RandomName()
	newLastName := util.RandomName()
	newEmail := util.RandomEmail()
	newPassword := util.RandomString(8)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		ID: user.ID,
		FirstName: sql.NullString{
			String: newFirstName,
			Valid:  true,
		},
		LastName: sql.NullString{
			String: newLastName,
			Valid:  true,
		},
		Email: sql.NullString{
			String: newEmail,
			Valid:  true,
		},
		Password: sql.NullString{
			String: newHashedPassword,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, user, updatedUser)
	require.Equal(t, newFirstName, updatedUser.FirstName)
	require.NotEqual(t, user.FirstName, updatedUser.FirstName)
	require.Equal(t, newLastName, updatedUser.LastName)
	require.NotEqual(t, user.LastName, updatedUser.LastName)
	require.Equal(t, newEmail, updatedUser.Email)
	require.NotEqual(t, user.Email, updatedUser.Email)
	require.Equal(t, newHashedPassword, updatedUser.Password)
	require.NotEqual(t, user.Password, updatedUser.Password)
}
