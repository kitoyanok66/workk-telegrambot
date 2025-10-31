package middleware

import (
	"sync"
)

var (
	jwtStore = make(map[string]string)
	jwtMutex sync.RWMutex
)

func SetJWT(userID string, token string) {
	jwtMutex.Lock()
	defer jwtMutex.Unlock()
	jwtStore[userID] = token
}

func GetJWT(userID string) (string, bool) {
	jwtMutex.RLock()
	defer jwtMutex.RUnlock()
	t, ok := jwtStore[userID]
	return t, ok
}
