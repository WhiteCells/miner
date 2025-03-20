package settlement

import (
	"sync"
)

var userLocks = make(map[int]*sync.Mutex)

func getUserLock(userID int) *sync.Mutex {
	if _, exists := userLocks[userID]; !exists {
		userLocks[userID] = &sync.Mutex{}
	}
	return userLocks[userID]
}
