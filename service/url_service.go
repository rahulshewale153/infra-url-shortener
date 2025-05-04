package service

import (
	"context"
	"errors"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/rahulshewale153/infra-url-shortener/repository"
	"github.com/rahulshewale153/infra-url-shortener/utils"
)

type urlService struct {
	urlStorageRepo repository.URLStorageRepo
}

func NewURLService(urlStorageRepo repository.URLStorageRepo) URLService {
	return &urlService{urlStorageRepo: urlStorageRepo}
}

// generate short url and store into the map
func (s *urlService) GetURLShortener(ctx context.Context, originalURL string) (string, error) {
	//validate the url
	u, err := url.Parse(originalURL)
	if err != nil || strings.TrimSpace(originalURL) == "" {
		log.Println("GetURLShortener: invalid url pass by user", originalURL)
		return "", errors.New("provided url is invalid")
	}

	//check the original url already exist in the map
	if originalURL, ok := s.urlStorageRepo.GetOriginalURL(ctx, originalURL); ok == nil {
		log.Println("GetURLShortener: url already exist in the map")
		return originalURL, nil
	}

	//get short url length by current millisecond value
	shortURLID := utils.GenerateEncodeBase62(int(time.Millisecond))

	//store the url data into map
	err = s.urlStorageRepo.Store(ctx, shortURLID, originalURL, strings.ToLower(u.Hostname()))
	if err != nil {
		log.Println("GetURLShortener: data couldn't be store")
		return "", errors.New("generate request couldn't be processed")
	}

	return shortURLID, nil
}
