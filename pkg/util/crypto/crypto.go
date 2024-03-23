package crypto

import (
	"bytes"
	"crypto/sha512"

	"golang.org/x/crypto/bcrypt"
)

func preparePlainTextToBeHashed(plainText string) (preparedText string) {
	hashedInput := sha512.Sum512_256([]byte(plainText))
	trimmedHash := bytes.Trim(hashedInput[:], "\x00")
	preparedText = string(trimmedHash)

	return preparedText
}

func CreateHash(plainText string) (hashText *string, err error) {
	preparedPlainText := preparePlainTextToBeHashed(plainText)
	hashTextInBytes, err := bcrypt.GenerateFromPassword([]byte(preparedPlainText), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	hashTextValue := string(hashTextInBytes)
	hashText = &hashTextValue

	return hashText, nil
}

func CompareHash(plainText string, hashText string) (textsMatch bool) {
	preparedPlainText := preparePlainTextToBeHashed(plainText)
	plainTextInBytes := []byte(preparedPlainText)
	hashTextInBytes := []byte(hashText)
	err := bcrypt.CompareHashAndPassword(hashTextInBytes, plainTextInBytes)
	textsMatch = err == nil

	return textsMatch
}
