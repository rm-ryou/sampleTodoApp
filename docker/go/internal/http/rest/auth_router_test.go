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

func TestUsersSignUp(t *testing.T) {
	tests := []struct {
		title      string
		reqDataStr string
		statusCode int
	}{
		{
			title: "valid request",
			reqDataStr: fmt.Sprintf(`{"name":"%s","email":"%s","password":"%s"}`,
				testdata.UserTestData[1].Name,
				testdata.UserTestData[1].Email,
				testdata.UserTestData[1].Password,
			),
			statusCode: http.StatusOK,
		},
		{
			title: "invalid request data",
			reqDataStr: fmt.Sprintf(`{"email":"%s","password":"%s"}`,
				testdata.UserTestData[1].Email,
				testdata.UserTestData[1].Password,
			),
			statusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			reqData := strings.NewReader(test.reqDataStr)
			url := fmt.Sprintf("%s/api/v1/auth/users/sign_up", baseURL)
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, url, reqData)

			router.ServeHTTP(res, req)

			assert.Equal(t, test.statusCode, res.Code)
		})
	}
}

func TestUsersSignIn(t *testing.T) {
	tests := []struct {
		title      string
		reqDataStr string
		statusCode int
	}{
		{
			title: "valid request",
			reqDataStr: fmt.Sprintf(`{"email":"%s","password":"%s"}`,
				testdata.UserTestData[1].Email,
				testdata.UserTestData[1].Password,
			),
			statusCode: http.StatusOK,
		},
		{
			title: "invalid request data",
			reqDataStr: fmt.Sprintf(`{"password":"%s"}`,
				testdata.UserTestData[1].Password,
			),
			statusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			reqData := strings.NewReader(test.reqDataStr)
			url := fmt.Sprintf("%s/api/v1/auth/users/sign_in", baseURL)
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, url, reqData)

			router.ServeHTTP(res, req)

			assert.Equal(t, test.statusCode, res.Code)
		})
	}
}

func TestAdminsSignIn(t *testing.T) {
	tests := []struct {
		title      string
		reqDataStr string
		statusCode int
	}{
		{
			title: "valid request",
			reqDataStr: fmt.Sprintf(`{"email":"%s","password":"%s"}`,
				testdata.UserTestData[0].Email,
				testdata.UserTestData[0].Password,
			),
			statusCode: http.StatusOK,
		},
		{
			title: "invalid request data",
			reqDataStr: fmt.Sprintf(`{"password":"%s"}`,
				testdata.UserTestData[0].Password,
			),
			statusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			reqData := strings.NewReader(test.reqDataStr)
			url := fmt.Sprintf("%s/api/v1/auth/admins/sign_in", baseURL)
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, url, reqData)

			router.ServeHTTP(res, req)

			assert.Equal(t, test.statusCode, res.Code)
		})
	}
}
