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

type UserResBody struct {
	Data []entity.User `json:"data"`
}

func TestGetUsers(t *testing.T) {
	url := fmt.Sprintf("%s/api/v1/users", baseURL)
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)

	router.ServeHTTP(res, req)

	var userResBody UserResBody
	if err := json.Unmarshal(res.Body.Bytes(), &userResBody); err != nil {
		t.Error(err)
	}
	users := userResBody.Data

	assert.Equal(t, 1, len(users))
	assert.Equal(t, testdata.UserTestData[1].ID, users[0].ID)
	assert.Equal(t, http.StatusOK, res.Code)
}
