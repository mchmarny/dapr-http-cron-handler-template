package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestRunOptions(t *testing.T) {
	t.Parallel()

	r := gin.Default()
	r.OPTIONS("/", optionsHandler)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodOptions, "/", nil)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRunHandler(t *testing.T) {
	t.Parallel()

	r := gin.Default()
	r.OPTIONS("/", runHandler)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodOptions, "/", nil)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
