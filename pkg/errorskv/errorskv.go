package errorskv

import "errors"

var (
	ErrKeyNotFound  = errors.New("key not found")
	ErrKeyExpired   = errors.New("key expired")
	ErrMaxKeysLimit = errors.New("max keys limit reached")
	ErrMaxMemLimit  = errors.New("max memory limit reached")
)
