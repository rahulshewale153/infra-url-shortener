package repository

import "context"

type URLStorageRepo interface {
	Store(ctx context.Context, shortURLID string, originalURL string, domain string) error
	GetOriginalURL(ctx context.Context, shortURLID string) (string, error)
	GetTop3Domain(ctx context.Context) ([]string, error)
}
