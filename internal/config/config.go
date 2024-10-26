package config

import "time"

type Config struct {
	CleanupInterval time.Duration
	MaxMemorySize   int64
	MaxKeys         int
}

func NewDefaultConfig() *Config {
	return &Config{
		CleanupInterval: time.Second,
		MaxMemorySize:   1 << 30, // 1GB
		MaxKeys:         1000000,
	}
}
