package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mock "github.com/rahulshewale153/infra-url-shortener/mock/service"
	"github.com/stretchr/testify/assert"
)

func TestURLShortener(t *testing.T) {
	mockURLService := new(mock.MockURLService)
	urlHandler := NewURLHandler(mockURLService)

	t.Run("invalid JSON request by user, function should be return an error", func(t *testing.T) {
		invalidRequest := `\invalid_json`
		req := httptest.NewRequest(http.MethodPost, "/url/shortener", strings.NewReader(invalidRequest))
		w := httptest.NewRecorder()

		urlHandler.URLShortener(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("url Shortener function couldn't process the request, function should be return an error", func(t *testing.T) {
		validRequest := `{"url": "http://www.originalurl.com/21324541"}`
		originalURL := "http://www.originalurl.com/21324541"
		expectedError := errors.New("generate request couldn't be processed")
		mockURLService.On("GetURLShortener", context.Background(), originalURL).Return("", expectedError).Once()

		req := httptest.NewRequest(http.MethodPost, "/url/shortener", strings.NewReader(validRequest))
		w := httptest.NewRecorder()

		urlHandler.URLShortener(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("valid request by user, function should be return short url", func(t *testing.T) {
		validRequest := `{"url": "http://www.originalurl.com/21324541"}`
		originalURL := "http://www.originalurl.com/21324541"
		expectedShortURL := "http://localhost:8080/Short12w"
		mockURLService.On("GetURLShortener", context.Background(), originalURL).Return("Short12w", nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/url/shortener", strings.NewReader(validRequest))
		w := httptest.NewRecorder()

		urlHandler.URLShortener(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), expectedShortURL)
	})

}
