package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := RandomString(12)

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
	require.NotEqual(t, password, hashedPassword)

	err = ValidatePassword(password, hashedPassword)
	require.NoError(t, err)

	wrongPassword := RandomString(8)
	err = ValidatePassword(wrongPassword, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedWrongPassword, err := HashPassword(wrongPassword)
	require.NoError(t, err)
	require.NotEmpty(t, hashedWrongPassword)
	require.NotEqual(t, hashedPassword, hashedWrongPassword)

}

func TestWrongPassword(t *testing.T) {
	password := RandomString(9)

	hashedPassword, err := HashPassword(RandomString(11))
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = ValidatePassword(password, hashedPassword)
	require.Error(t, err)
}
