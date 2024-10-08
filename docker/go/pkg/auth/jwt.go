package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rm-ryou/sampleTodoApp/pkg/utils"
)

var mySigningKey []byte

type CustomClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func InitializeSigningKey(key string) {
	mySigningKey = []byte(key)
}

func GenerateToken(userId int, timer utils.Timer) (string, error) {
	claims := CustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(timer.Now().Add(3 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(timer.Now()),
			NotBefore: jwt.NewNumericDate(timer.Now()),
			Issuer:    "sampleTodoApp",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func VerifyToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("unknown claims type, cannot proceed")
	}

	return claims, nil
}
