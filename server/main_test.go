package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWsHandler(t *testing.T) {

}

func TestRootHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(rootHandler)
	handler.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "OK status code is expected")
	assert.Equal(t, "Works", response.Body.String(), "Incorrect string found")
}

func TestWaitForRequests(t *testing.T) {

}

func TestReplaceQuestionMarks(t *testing.T) {
	tables := []struct{ input, output string }{{"ima", "ima"}, {"?sa?", "!sa!"}}
	for _, table := range tables {
		result := replaceQuestionMarks(table.input)
		assert.Equal(t, result, table.output, "Question marks should be replaced with exclamation marks")
	}
}

func TestContains(t *testing.T) {

}
