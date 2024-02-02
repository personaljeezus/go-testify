package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	if assert.Len(t, list, totalCount) {
		assert.Equal(t, assert.Len, totalCount)
	}
}
func TestIsYourReqCorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	require.NotEmpty(t, req)
	status := responseRecorder.Code
	assert.Equal(t, status, http.StatusOK)
}
func TestIfCityCorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	actualCity := req.URL.Query().Get("city")
	expected := "moscow"
	assert.Equal(t, actualCity, expected)

}
