package memory

import (
	"sync"

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
