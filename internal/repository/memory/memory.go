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
	go db.cleanupLoop()
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

func (db *MemoryDB) Delete(key string) error {

	db.mu.Lock()
	defer db.mu.Unlock()

	if !db.Exists(key) {
		return errorskv.ErrKeyNotFound
	}

	delete(db.data, key)
	db.logger.Info.Println("Deleted key: " + colors.Red + key + colors.Reset)
	return errorskv.ErrSuccessDelete

}

func (db *MemoryDB) Exists(key string) bool {

	value, exists := db.data[key]
	if !exists {
		return false
	}

	if !value.ExpiresAt.IsZero() && time.Now().After(value.ExpiresAt) {
		delete(db.data, key)
		return false
	}

	return true
}

func (db *MemoryDB) SetTTL(key string, ttl time.Duration) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	value, exists := db.data[key]
	if !exists {
		return errorskv.ErrKeyNotFound
	}

	value.ExpiresAt = time.Now().Add(ttl)
	db.data[key] = value

	db.logger.Info.Printf("Updated TTL for key: %s", key)
	return nil
}
func (db *MemoryDB) GetTTL(key string) (time.Duration, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	value, exists := db.data[key]
	if !exists {
		return 0, errorskv.ErrKeyNotFound
	}

	if value.ExpiresAt.IsZero() {
		return 0, nil
	}

	ttl := time.Until(value.ExpiresAt)
	if ttl < 0 {
		delete(db.data, key)
		return 0, errorskv.ErrKeyExpired
	}

	return ttl, nil
}

// clean up expired keys
func (db *MemoryDB) cleanupLoop() {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		db.mu.Lock()
		for k, v := range db.data {
			if !v.ExpiresAt.IsZero() && time.Now().After(v.ExpiresAt) {
				delete(db.data, k)
				db.logger.Info.Printf("Expired key removed: %s", k)
			}
		}
		db.mu.Unlock()
	}
}
