package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	mock "github.com/rahulshewale153/infra-url-shortener/mock/service"
	"github.com/stretchr/testify/assert"
	mocks "github.com/stretchr/testify/mock"
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

func TestGetOriginalURL(t *testing.T) {
	mockURLService := new(mock.MockURLService)
	urlHandler := NewURLHandler(mockURLService)

	t.Run("invalid short URL ID, function should be return an error", func(t *testing.T) {
		shortURLID := "invalidShortID"
		mockURLService.On("GetOriginalURL", mocks.Anything, shortURLID).Return("", errors.New("original url not found")).Once()
		req := httptest.NewRequest(http.MethodGet, "/"+shortURLID, nil)
		req = mux.SetURLVars(req, map[string]string{"short_url_id": shortURLID})
		w := httptest.NewRecorder()

		urlHandler.GetOriginalURL(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("valid short URL ID, function should be return original URL", func(t *testing.T) {
		shortURLID := "Short12w"
		expectedOriginalURL := "http://www.originalurl.com/21324541"
		mockURLService.On("GetOriginalURL", mocks.Anything, shortURLID).Return(expectedOriginalURL, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/"+shortURLID, nil)
		req = mux.SetURLVars(req, map[string]string{"short_url_id": shortURLID})
		w := httptest.NewRecorder()

		urlHandler.GetOriginalURL(w, req)
		assert.Equal(t, http.StatusFound, w.Code)
		assert.Equal(t, expectedOriginalURL, w.Header().Get("Location"))
	})
}
