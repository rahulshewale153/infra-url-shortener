package repository

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockURLService struct {
	mock.Mock
}

func (m *MockURLService) GetURLShortener(ctx context.Context, originalURL string) (string, error) {
	args := m.Called(ctx, originalURL)
	return args.String(0), args.Error(1)
}

func (m *MockURLService) GetOriginalURL(ctx context.Context, shortURLID string) (string, error) {
	args := m.Called(ctx, shortURLID)
	return args.String(0), args.Error(1)
}

func (m *MockURLService) GetTop3Domain(ctx context.Context) ([]string, error) {
	args := m.Called(ctx)
	return args.Get(0).([]string), args.Error(1)
}
