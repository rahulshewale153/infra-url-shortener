package service

import (
	"context"
	"errors"
	"testing"

	mock "github.com/rahulshewale153/infra-url-shortener/mock/repository"
	"github.com/stretchr/testify/assert"
	mocks "github.com/stretchr/testify/mock"
)

func TestGetURLShortener(t *testing.T) {
	mockUrlStoreRepo := new(mock.MockURLStore)
	urlService := NewURLService(mockUrlStoreRepo)

	t.Run("invalid url send by user, function should be return an error", func(t *testing.T) {
		invalidURL := "http://%invalid_url"
		expectedError := errors.New("provided url is invalid")
		_, err := urlService.GetURLShortener(context.Background(), invalidURL)
		assert.Equal(t, expectedError, err)
	})

	t.Run("empty url send by user, function should be return an error", func(t *testing.T) {
		emptyURL := ""
		expectedError := errors.New("provided url is invalid")
		_, err := urlService.GetURLShortener(context.Background(), emptyURL)
		assert.Equal(t, expectedError, err)
	})

	t.Run("url already exist in the map, function should be return short url", func(t *testing.T) {
		validURL := "http://www.originalurl.com/21324541"
		mockUrlStoreRepo.On("GetOriginalURL", context.Background(), validURL).Return(validURL, nil).Once()
		shortURL, err := urlService.GetURLShortener(context.Background(), validURL)
		assert.NoError(t, err)
		assert.Equal(t, validURL, shortURL)
	})

	t.Run("store error occurred, function should be return an error", func(t *testing.T) {
		validURL := "http://www.originalurl.com/21324541"
		expectedError := errors.New("generate request couldn't be processed")
		mockUrlStoreRepo.On("GetOriginalURL", context.Background(), validURL).Return("", errors.New("short URL not found")).Once()
		mockUrlStoreRepo.On("Store", context.Background(), mocks.Anything, validURL, "www.originalurl.com").Return(expectedError).Once()

		shortURL, err := urlService.GetURLShortener(context.Background(), validURL)
		assert.Equal(t, expectedError, err)
		assert.Empty(t, shortURL)
	})

	t.Run("valid url send by user, function should be return short url", func(t *testing.T) {
		validURL := "http://www.originalurl.com/21324541"
		mockUrlStoreRepo.On("GetOriginalURL", context.Background(), validURL).Return("", errors.New("short URL not found")).Once()
		mockUrlStoreRepo.On("Store", context.Background(), mocks.Anything, validURL, "www.originalurl.com").Return(nil).Once()

		shortURL, err := urlService.GetURLShortener(context.Background(), validURL)
		assert.NoError(t, err)
		assert.NotEmpty(t, shortURL)
	})

}

func TestGetOriginalURL(t *testing.T) {
	mockUrlStoreRepo := new(mock.MockURLStore)
	urlService := NewURLService(mockUrlStoreRepo)

	t.Run("error occurred while retrieving original url, function should be return an error", func(t *testing.T) {
		validShortURL := "short123"
		expectedError := errors.New("invalid short url")
		mockUrlStoreRepo.On("GetOriginalURL", context.Background(), validShortURL).Return("", expectedError).Once()

		originalURL, err := urlService.GetOriginalURL(context.Background(), validShortURL)
		assert.Equal(t, expectedError, err)
		assert.Empty(t, originalURL)
	})

	t.Run("valid short url send by user, function should be return original url", func(t *testing.T) {
		validShortURL := "short123"
		expectedOriginalURL := "http://www.originalurl.com/21324541"
		mockUrlStoreRepo.On("GetOriginalURL", context.Background(), validShortURL).Return(expectedOriginalURL, nil).Once()

		originalURL, err := urlService.GetOriginalURL(context.Background(), validShortURL)
		assert.NoError(t, err)
		assert.Equal(t, expectedOriginalURL, originalURL)
	})

}

func TestTop3Domain(t *testing.T) {
	mockUrlStoreRepo := new(mock.MockURLStore)
	urlService := NewURLService(mockUrlStoreRepo)

	t.Run("error occurred while retrieving top  domain, function should be return an error", func(t *testing.T) {
		expectedError := errors.New("top 3 domain couldn't be retrieve")
		mockUrlStoreRepo.On("GetTop3Domain", context.Background()).Return([]string{}, expectedError).Once()

		topDomains, err := urlService.GetTop3Domain(context.Background())
		assert.Equal(t, expectedError, err)
		assert.Empty(t, topDomains)
	})

	t.Run("valid request by user, function should be return top 3 domain", func(t *testing.T) {
		expectedTopDomains := []string{"www.originalurl.com", "www.example.com", "www.test.com"}
		mockUrlStoreRepo.On("GetTop3Domain", context.Background()).Return(expectedTopDomains, nil).Once()

		topDomains, err := urlService.GetTop3Domain(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, expectedTopDomains, topDomains)
	})
}
