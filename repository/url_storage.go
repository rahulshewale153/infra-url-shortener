package repository

import (
	"context"
	"errors"
	"sync"
)

type urlStorageRepo struct {
	urlCollection      map[string]string
	shortUrlCollection map[string]string
	domainRequest      map[string]int
	mu                 sync.RWMutex
}

func NewURLStorageRepo() URLStorageRepo {
	//initialize map
	return &urlStorageRepo{
		urlCollection:      make(map[string]string),
		shortUrlCollection: make(map[string]string),
		domainRequest:      make(map[string]int),
	}
}

// store the url data into map
func (r *urlStorageRepo) Store(ctx context.Context, shortURLID string, originalURL string, domain string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.urlCollection[originalURL] = shortURLID
	r.shortUrlCollection[shortURLID] = originalURL
	r.domainRequest[domain]++

	return nil
}

// Get the original url by short url
func (r *urlStorageRepo) GetOriginalURL(ctx context.Context, shortURLID string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	originalURL, ok := r.shortUrlCollection[shortURLID]
	if !ok {
		return "", errors.New("short URL not found")
	}

	return originalURL, nil
}
