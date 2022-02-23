package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	token, err := New(Data{1})
	assert.NoError(t, err)
	t.Log(token)
	userID, err := GetUserIDFromToken(token)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), userID)
}

func TestIsValid(t *testing.T) {
	ok, err := IsValid("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
		"eyJleHAiOjE2NDU2Njk3NzUsIm5iZiI6MTY0NTU4MzM3NSwiaWF0IjoxNjQ1NTgzMzc1LCJ1c2VyX2lkIjoxfQ.N3D2SRS207JT3FofkfWaIHP4MQ04CClfFi2BzyA6Eq8")
	assert.NoError(t, err)
	assert.True(t, ok)
}
