package memory

import (
	"sync"
	"time"

	"github.com/Dostonlv/kv/internal/domain/models"
	"github.com/Dostonlv/kv/internal/logger"
	"github.com/Dostonlv/kv/pkg/colors"
	"github.com/Dostonlv/kv/pkg/errorskv"
)

type MemoryDB struct {
	mu     sync.RWMutex
	data   map[string]models.Value
	logger *logger.Logger
}

func NewMemoryDB(logger *logger.Logger) *MemoryDB {
	db := &MemoryDB{
		data:   make(map[string]models.Value),
		logger: logger,
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
	db.logger.Info.Println("Set key:" + colors.Green + key + colors.Reset)
	return nil
}

func (db *MemoryDB) Get(key string) (interface{}, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	value, exists := db.data[key]
	if !exists {
		return nil, errorskv.ErrKeyNotFound
	}
	if !value.ExpiresAt.IsZero() && time.Now().After(value.ExpiresAt) {

		db.mu.RUnlock()
		db.mu.Lock()
		delete(db.data, key)
		db.mu.Unlock()
		db.mu.RLock()

		return nil, errorskv.ErrKeyExpired
	}

	return value.Data, nil
}
