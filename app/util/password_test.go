package util

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	hashed1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashed1)

	err = CheckPassword(password, hashed1)
	require.NoError(t, err)

	err = CheckPassword("wrong_password", hashed1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashed2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashed1)
	require.NotEqual(t, hashed1, hashed2)

}
