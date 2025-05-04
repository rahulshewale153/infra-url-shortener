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
