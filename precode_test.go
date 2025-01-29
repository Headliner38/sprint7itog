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
	//totalCount := 4
	req, err := http.NewRequest("GET", "/cafe?count=5&city=moscow", nil) // здесь нужно создать запрос к сервису
	require.NoError(t, err)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "expected status code 200")

	kolvoBody := responseRecorder.Body.String()
	list := strings.Split(kolvoBody, ",")
	assert.Len(t, list, 4, "expected 4 cities")
}

func TestWrongCityValue(t *testing.T) {
	//cityValue := "penza"

	req, err := http.NewRequest("GET", "/cafe?count=2&city=penza", nil)
	require.NoError(t, err)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "expected status code 400")

	c := "wrong city value"
	assert.Equal(t, c, responseRecorder.Body.String(), "expected body error")

}

func TestMainHandlerWhenOk(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	require.NoError(t, err)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "expected status code 200")

	assert.NotEmpty(t, responseRecorder.Body.String(), "expected not empty body")
}
