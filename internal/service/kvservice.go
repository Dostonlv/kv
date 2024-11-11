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

func (s *KVService) Delete(key string) error {
	return s.repo.Delete(key)
}

func (s *KVService) Exists(key string) bool {
	return s.repo.Exists(key)
}

func (s *KVService) SetTTL(key string, ttl time.Duration) error {
	return s.repo.SetTTL(key, ttl)
}

func (s *KVService) GetTTL(key string) (time.Duration, error) {
	return s.repo.GetTTL(key)
}
