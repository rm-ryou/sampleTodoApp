package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rm-ryou/sampleTodoApp/internal/entity"
	testdata "github.com/rm-ryou/sampleTodoApp/test/data"
	"github.com/stretchr/testify/assert"
)

func TestUsersSignUp(t *testing.T) {
	expectedUser := testdata.UserTestData[1]
	reqDataStr := fmt.Sprintf(`{"name":"%s","email":"%s","password":"%s"}`,
		expectedUser.Name,
		expectedUser.Email,
		expectedUser.Password)

	reqData := strings.NewReader(reqDataStr)
	url := fmt.Sprintf("%s/api/v1/auth/users/sign_up", baseURL)
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, url, reqData)

	router.ServeHTTP(res, req)

	var userResponse UserResponse[entity.User]
	if err := json.Unmarshal(res.Body.Bytes(), &userResponse); err != nil {
		t.Error(err)
	}
	user := userResponse.Data

	assert.Equal(t, expectedUser.Name, user.Name)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestUsersSignIn(t *testing.T) {
	tests := []struct {
		title      string
		body       string
		statusCode int
	}{
		{
			title: "valid request",
			body: fmt.Sprintf(`{"email":"%s","password":"%s"}`,
				testdata.UserTestData[1].Email,
				testdata.UserTestData[1].Password),
			statusCode: http.StatusOK,
		},
		{
			title: "invalid request",
			body: fmt.Sprintf(`{"email":"%s"}`,
				testdata.UserTestData[1].Email),
			statusCode: http.StatusBadRequest,
		},
		{
			title: "admin can't sign in as user",
			body: fmt.Sprintf(`{"email":"%s","password":"%s"}`,
				testdata.UserTestData[0].Email,
				testdata.UserTestData[0].Password),
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

func TestAdminsSignIn(t *testing.T) {
	tests := []struct {
		title      string
		body       string
		statusCode int
	}{
		{
			title: "valid request",
			body: fmt.Sprintf(`{"email":"%s","password":"%s"}`,
				testdata.UserTestData[0].Email,
				testdata.UserTestData[0].Password),
			statusCode: http.StatusOK,
		},
		{
			title: "invalid request",
			body: fmt.Sprintf(`{"email":"%s"}`,
				testdata.UserTestData[0].Email),
			statusCode: http.StatusBadRequest,
		},
		{
			title: "user can't sign in as admin",
			body: fmt.Sprintf(`{"email":"%s","password":"%s"}`,
				testdata.UserTestData[1].Email,
				testdata.UserTestData[1].Password),
			statusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			reqData := strings.NewReader(test.body)
			url := fmt.Sprintf("%s/api/v1/auth/admins/sign_in", baseURL)
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, url, reqData)

			router.ServeHTTP(res, req)

			assert.Equal(t, test.statusCode, res.Code)
		})
	}
}
