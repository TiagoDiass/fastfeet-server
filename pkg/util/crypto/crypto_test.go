package crypto

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldCreateHashCorrectly(t *testing.T) {
	plainText := "password"
	hashText, err := CreateHash(plainText)

	require.NoError(t, err)
	require.NotEmpty(t, &hashText)
}

func TestShouldCompareHashCorrectly(t *testing.T) {
	plainText := "password"
	hashText, err := CreateHash(plainText)

	require.NoError(t, err)

	passwordsMatch := CompareHash(plainText, &hashText)

	require.Equal(t, passwordsMatch, true)
}

func TestShouldCompareHashWithStringsNotMatching(t *testing.T) {
	plainText := "password"
	hashText, _ := CreateHash(plainText)
	incorrectPlainText := "wrongpassword"

	passwordsMatch := CompareHash(incorrectPlainText, &hashText)

	require.Equal(t, passwordsMatch, false)
}
