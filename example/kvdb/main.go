package main

import (
	"log"
	"time"

	// "github.com/Dostonlv/kv/internal/config"
	"github.com/Dostonlv/kv/internal/logger"
	"github.com/Dostonlv/kv/internal/repository/memory"
	"github.com/Dostonlv/kv/internal/service"
)

func main() {
	// cfg := config.NewDefaultConfig()
	logger := logger.NewLogger()

	// Create repository
	db := memory.NewMemoryDB(logger)

	// Create service
	kvService := service.NewKVService(db)

	// Use the service
	if err := kvService.Set("key", "kv", time.Minute); err != nil {
		log.Fatal(err)
	}

	// Set a key with a value and a TTL
	if err := kvService.SetTTL("key", time.Minute); err != nil {
		log.Fatal(err)
	}

	value, err := kvService.Get("bb")
	if err != nil {
		log.Fatal(err)
	}

	err = kvService.Delete("aa")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(value)
}
