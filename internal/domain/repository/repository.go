package repository

import (
	"time"
)

type KVRepository interface {
	Set(key string, value interface{}, ttl time.Duration) error
	Get(key string) (interface{}, error)
	Delete(key string) error
	Exists(key string) bool
	Keys() []string
	Size() int
	Clear() error
	SetTTL(key string, ttl time.Duration) error
	GetTTL(key string) (time.Duration, error)
}
