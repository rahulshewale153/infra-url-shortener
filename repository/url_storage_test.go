package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	repo := NewURLStorageRepo()
	shortID := "Short12w"
	originalURL := "http://www.originalurl.com/21324541"
	domain := "www.originalurl.co"
	t.Run("Store Short and original url", func(t *testing.T) {
		err := repo.Store(context.Background(), shortID, originalURL, domain)
		assert.Empty(t, err)
	})
}

func TestGetOriginalURL(t *testing.T) {
	repo := NewURLStorageRepo()
	shortID := "Short12w"

	t.Run("Original url not found,function should be return an error ", func(t *testing.T) {
		originalURL, err := repo.GetOriginalURL(context.Background(), shortID)
		assert.NotEmpty(t, err)
		assert.Empty(t, originalURL)
	})

	t.Run("original url found in collection", func(t *testing.T) {
		// Store the original URL first
		err := repo.Store(context.Background(), shortID, "http://www.originalurl.com/21324541", "www.originalurl.co")
		assert.Empty(t, err)

		// Now test the retrieval of the original URL
		originalURL, err := repo.GetOriginalURL(context.Background(), shortID)
		assert.Empty(t, err)
		assert.Equal(t, originalURL, "http://www.originalurl.com/21324541")
	})
}
