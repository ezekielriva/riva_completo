package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingController(t *testing.T) {
	router := SetupRouter(nil)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	body := struct {
		Db struct {
			OpenConnections int `json:"OpenConnections"`
		} `json:"db,omitempty"`
		Status string
	}{}

	err := json.Unmarshal(w.Body.Bytes(), &body)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", body.Status)
	assert.Equal(t, 0, body.Db.OpenConnections)
}
