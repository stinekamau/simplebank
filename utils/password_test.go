package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)
	hpassword, err := HashPassword(password)

	require.NotEmpty(t, hpassword)
	require.NoError(t, err)

	err = CheckPassword(password, hpassword)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hpassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())




}
