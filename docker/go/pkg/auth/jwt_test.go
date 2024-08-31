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

func TestVerifyToken(t *testing.T) {
	user := testdata.UserTestData[1]
	t.Run("valid token", func(t *testing.T) {
		token, err := GenerateToken(user.ID, utils.RealTime{})
		if err != nil {
			t.Error(err)
		}

		claims, err := VerifyToken(token)

		assert.Nil(t, err)
		assert.Equal(t, claims.UserID, user.ID)
	})

	t.Run("Invalid token", func(t *testing.T) {
		token := "hoge"

		claims, err := VerifyToken(token)

		assert.NotNil(t, err)
		assert.Nil(t, claims)
	})

	t.Run("Expired token", func(t *testing.T) {
		baseTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
		mt := utils.NewMockTime(baseTime)
		token, err := GenerateToken(user.ID, mt)
		if err != nil {
			t.Error(err)
		}

		claims, err := VerifyToken(token)

		assert.NotNil(t, err)
		assert.Nil(t, claims)
	})
}
