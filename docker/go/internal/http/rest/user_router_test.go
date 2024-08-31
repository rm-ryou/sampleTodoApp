package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
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

	t.Run("parameter is number", func(t *testing.T) {
		expectedUser := testdata.UserTestData[1]
		url := fmt.Sprintf("%s/api/v1/users/%d", baseURL, expectedUser.ID)
		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, url, nil)

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
		url := fmt.Sprintf("%s/api/v1/users/hoge", baseURL)
		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, url, nil)

		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
}

func TestGetUsers(t *testing.T) {
	url := fmt.Sprintf("%s/api/v1/users", baseURL)
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)

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
