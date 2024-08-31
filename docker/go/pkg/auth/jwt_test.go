package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rm-ryou/sampleTodoApp/pkg/utils"
	testdata "github.com/rm-ryou/sampleTodoApp/test/data"
	"github.com/stretchr/testify/assert"
)

var SigningKey = "test"

func TestMain(m *testing.M) {
	InitializeSigningKey(SigningKey)

	m.Run()
}

func TestGenerateToken(t *testing.T) {
	user := testdata.UserTestData[1]
	baseTime := time.Now()
	mt := utils.NewMockTime(baseTime)

	tokenStr, err := GenerateToken(user.ID, mt)
	if err != nil {
		t.Fatal(err)
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SigningKey), nil
	})
	if err != nil {
		t.Fatal(err)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		t.Fatal("Failed to parse claims")
	}

	expectedTime := jwt.NewNumericDate(mt.Now().Add(3 * time.Hour))

	assert.True(t, token.Valid)
	assert.Equal(t, user.ID, claims.UserID)
	assert.Equal(t, expectedTime, claims.ExpiresAt)
}
