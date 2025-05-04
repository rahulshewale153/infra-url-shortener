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
