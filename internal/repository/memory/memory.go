package memory

import (
	"sync"
	"time"

	"github.com/Dostonlv/kv/internal/domain/models"
)

type MemoryDB struct {
	mu   sync.RWMutex
	data map[string]models.Value
}

func NewMemoryDB() *MemoryDB {
	db := &MemoryDB{
		data: make(map[string]models.Value),
	}
	return db
}

func (db *MemoryDB) Set(key string, value interface{}, ttl time.Duration) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	expiresAt := time.Time{}
	if ttl > 0 {
		expiresAt = time.Now().Add(ttl)
	}

	db.data[key] = models.Value{
		Data:      value,
		ExpiresAt: expiresAt,
	}
	return nil
}
