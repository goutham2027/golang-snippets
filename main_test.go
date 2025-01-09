package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Hello(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rw := httptest.NewRecorder()
	hello(rw, req)
	resp := rw.Result()
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Status code should be 200, got %d", resp.StatusCode)
	}

	body, err := getResponseBody(resp)
	assert.Nil(t, err)

	assert.Equal(t, "Hello World", body)
}

func getResponseBody(resp *http.Response) (string, error) {
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
