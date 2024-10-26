package models

import "time"

type Value struct {
	Data      interface{}
	ExpiresAt time.Time
}
