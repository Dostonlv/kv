package config

import "time"

type Config struct {
	CleanupInterval time.Duration
	MaxMemorySize   int64
	MaxKeys         int
}
