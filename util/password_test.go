package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := RandomString(10)
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(hashedPassword, password)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(hashedPassword, wrongPassword)
	require.EqualError(t, err, "invalid password: crypto/bcrypt: hashedPassword is not the hash of the given password")
}
