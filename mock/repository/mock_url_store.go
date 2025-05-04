package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockURLStore struct {
	mock.Mock
}

func (m *MockURLStore) Store(ctx context.Context, shortURLID string, originalURL string, domain string) error {
	args := m.Called(ctx, shortURLID, originalURL, domain)
	return args.Error(0)
}

func (m *MockURLStore) GetOriginalURL(ctx context.Context, shortURLID string) (string, error) {
	args := m.Called(ctx, shortURLID)
	return args.String(0), args.Error(1)
}

func (m *MockURLStore) GetTop3Domain(ctx context.Context) ([]string, error) {
	args := m.Called(ctx)
	return args.Get(0).([]string), args.Error(1)
}
