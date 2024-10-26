package main

import (
	"log"
	"time"

	// "github.com/Dostonlv/kv/internal/config"
	"github.com/Dostonlv/kv/internal/repository/memory"
	"github.com/Dostonlv/kv/internal/service"
)

func main() {
	// cfg := config.NewDefaultConfig()

	// Create repository
	db := memory.NewMemoryDB()

	// Create service
	kvService := service.NewKVService(db)

	// Use the service
	if err := kvService.Set("key", "kv", time.Minute); err != nil {
		log.Fatal(err)
	}

	value, err := kvService.Get("key")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(value)
}
