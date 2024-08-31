package rest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	testdata "github.com/rm-ryou/sampleTodoApp/test/data"
	"github.com/stretchr/testify/assert"
)

func TestUsersSignIn(t *testing.T) {
	user := testdata.UserTestData[1]
	tests := []struct {
		title      string
		body       string
		statusCode int
	}{
		{
			title:      "valid request",
			body:       fmt.Sprintf(`{"email":"%s","password":"%s"}`, user.Email, user.Password),
			statusCode: http.StatusOK,
		},
		{
			title:      "invalid request",
			body:       fmt.Sprintf(`{"email":"%s"}`, user.Email),
			statusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			reqData := strings.NewReader(test.body)
			url := fmt.Sprintf("%s/api/v1/auth/users/sign_in", baseURL)
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, url, reqData)

			router.ServeHTTP(res, req)

			assert.Equal(t, test.statusCode, res.Code)
		})
	}
}
