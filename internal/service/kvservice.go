package service

import (
	"time"

	"github.com/Dostonlv/kv/internal/domain/repository"
)

type KVService struct {
	repo repository.KVRepository
}

func NewKVService(repo repository.KVRepository) *KVService {
	return &KVService{
		repo: repo,
	}
}

func (s *KVService) Set(key string, value interface{}, ttl time.Duration) error {
	return s.repo.Set(key, value, ttl)
}

func (s *KVService) Get(key string) (interface{}, error) {
	return s.repo.Get(key)
}
