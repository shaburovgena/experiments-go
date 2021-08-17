package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleAbout(t *testing.T) {
	s := New(newConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "health", nil)
	s.handler.HandleHealth(rec, req)
	assert.Equal(t, rec.Body.String(), "OK")
}
