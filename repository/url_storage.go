package repository

import (
	"context"
	"errors"
	"sort"
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

// Get Top 3 Domain
func (r *urlStorageRepo) GetTop3Domain(ctx context.Context) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	type domainCount struct {
		domain string
		count  int
	}

	domainCounts := make([]domainCount, 0, len(r.domainRequest))
	for domain, count := range r.domainRequest {
		domainCounts = append(domainCounts, domainCount{domain: domain, count: count})
	}

	// Sort the domain counts in descending order
	sort.Slice(domainCounts, func(i, j int) bool {
		return domainCounts[i].count > domainCounts[j].count
	})

	// Get the top 3 domains
	top3Domains := make([]string, 0, 3)
	for i := 0; i < 3 && i < len(domainCounts); i++ {
		top3Domains = append(top3Domains, domainCounts[i].domain)
	}

	return top3Domains, nil
}
