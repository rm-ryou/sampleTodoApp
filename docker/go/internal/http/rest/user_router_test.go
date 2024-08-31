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

type UserResponser interface {
	[]entity.User | entity.User
}

type UserResponse[T UserResponser] struct {
	Data T `json:"data"`
}

func TestGetUser(t *testing.T) {
	t.Run("parameter is valid", func(t *testing.T) {
		expectedUser := testdata.UserTestData[1]
		url := fmt.Sprintf("%s/api/v1/users/%d", baseURL, expectedUser.ID)
		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, url, nil)

		if err := setHeader(expectedUser.ID, req); err != nil {
			t.Error(err)
		}

		router.ServeHTTP(res, req)

		var userResponse UserResponse[entity.User]
		if err := json.Unmarshal(res.Body.Bytes(), &userResponse); err != nil {
			t.Error(err)
		}
		user := userResponse.Data

		assert.Equal(t, expectedUser.ID, user.ID)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("parameter is not number", func(t *testing.T) {
		user := testdata.UserTestData[1]
		url := fmt.Sprintf("%s/api/v1/users/hoge", baseURL)
		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, url, nil)

		if err := setHeader(user.ID, req); err != nil {
			t.Error(err)
		}

		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
}

func TestGetUsers(t *testing.T) {
	user := testdata.UserTestData[1]
	url := fmt.Sprintf("%s/api/v1/users", baseURL)
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)

	if err := setHeader(user.ID, req); err != nil {
		t.Error(err)
	}

	router.ServeHTTP(res, req)

	var userResponse UserResponse[[]entity.User]
	if err := json.Unmarshal(res.Body.Bytes(), &userResponse); err != nil {
		t.Error(err)
	}
	users := userResponse.Data

	assert.Equal(t, 1, len(users))
	assert.Equal(t, testdata.UserTestData[1].ID, users[0].ID)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestEditUser(t *testing.T) {
	t.Run("update user name", func(t *testing.T) {
		subjectUser := testdata.UserTestData[1]

		updateName := "Update Name"
		reqDataStr := fmt.Sprintf(`{"name":"%s"}`, updateName)

		reqData := strings.NewReader(reqDataStr)
		url := fmt.Sprintf("%s/api/v1/users/%d", baseURL, subjectUser.ID)
		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPatch, url, reqData)

		if err := setHeader(subjectUser.ID, req); err != nil {
			t.Error(err)
		}

		router.ServeHTTP(res, req)

		var userResponse UserResponse[entity.User]
		if err := json.Unmarshal(res.Body.Bytes(), &userResponse); err != nil {
			t.Error(err)
		}
		user := userResponse.Data

		assert.Equal(t, updateName, user.Name)
		assert.Equal(t, subjectUser.Email, user.Email)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("parameter is not number", func(t *testing.T) {
		user := testdata.UserTestData[1]
		url := fmt.Sprintf("%s/api/v1/users/hoge", baseURL)
		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPatch, url, nil)

		if err := setHeader(user.ID, req); err != nil {
			t.Error(err)
		}

		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		title      string
		param      string
		statusCode int
	}{
		{title: "parameter is valid", param: "1", statusCode: http.StatusOK},
		{title: "parameter is not valid", param: "hoge", statusCode: http.StatusBadRequest},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			user := testdata.UserTestData[1]
			url := fmt.Sprintf("%s/api/v1/users/%s", baseURL, test.param)
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, url, nil)

			if err := setHeader(user.ID, req); err != nil {
				t.Error(err)
			}

			router.ServeHTTP(res, req)

			assert.Equal(t, test.statusCode, res.Code)
		})
	}
}
