package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rm-ryou/sampleTodoApp/pkg/auth"
	"github.com/rm-ryou/sampleTodoApp/pkg/utils"
	testdata "github.com/rm-ryou/sampleTodoApp/test/data"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	user := testdata.UserTestData[1]

	tests := []struct {
		title      string
		setHeader  func(r *http.Request)
		statusCode int
	}{
		{
			title: "valid token",
			setHeader: func(r *http.Request) {
				token, err := auth.GenerateToken(user.ID, utils.RealTime{})
				if err != nil {
					t.Fatal(err)
				}

				r.Header.Set("Authorization", "Bearer "+token)
			},
			statusCode: http.StatusOK,
		},
		{
			title:      "No token",
			setHeader:  func(r *http.Request) {},
			statusCode: http.StatusBadRequest,
		},
		{
			title: "Invalid header",
			setHeader: func(r *http.Request) {
				r.Header.Set("Authorization", "Bearer Invalid")
			},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			url := fmt.Sprintf("%s/auth", baseURL)
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, url, nil)

			test.setHeader(req)
			router.ServeHTTP(res, req)

			assert.Equal(t, test.statusCode, res.Code)
		})
	}
}
