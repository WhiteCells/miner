package settlement

import (
	"sync"
)

var userLocks = make(map[string]*sync.Mutex)

func getUserLock(userID string) *sync.Mutex {
	if _, exists := userLocks[userID]; !exists {
		userLocks[userID] = &sync.Mutex{}
	}
	return userLocks[userID]
}
