package service

import "context"

type URLService interface {
	GetURLShortener(ctx context.Context, originalURL string) (string, error)
	GetOriginalURL(ctx context.Context, shortURLID string) (string, error)
}
