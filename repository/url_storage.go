package repository

import (
	"context"
	"sync"
)

type urlStorageRepo struct {
	urlCollection      map[string]string
	shortUrlCollection map[string]string
	domainRequest      map[string]int
	mu                 sync.RWMutex
}

func NewURLStorageRepo() URLStorageRepo {
	return &urlStorageRepo{
		urlCollection:      make(map[string]string),
		shortUrlCollection: make(map[string]string),
		domainRequest:      make(map[string]int),
	}
}

func (r *urlStorageRepo) Store(ctx context.Context, shortURLID string, originalURL string, domain string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.urlCollection[originalURL] = shortURLID
	r.shortUrlCollection[shortURLID] = originalURL
	r.domainRequest[domain]++

	return nil
}
