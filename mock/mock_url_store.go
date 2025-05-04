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
