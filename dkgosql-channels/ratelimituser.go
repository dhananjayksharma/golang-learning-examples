package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type UserLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.Mutex
}

func NewUserLimiter() *UserLimiter {
	return &UserLimiter{
		limiters: make(map[string]*rate.Limiter),
	}
}

func (u *UserLimiter) getLimiter(userID string) *rate.Limiter {
	u.mu.Lock()
	defer u.mu.Unlock()

	limiter, exists := u.limiters[userID]
	if !exists {
		// 3 requests per second, burst = 3
		limiter = rate.NewLimiter(3, 3)
		u.limiters[userID] = limiter
	}
	return limiter
}

func (u *UserLimiter) Allow(userID string) bool {
	limiter := u.getLimiter(userID)
	return limiter.Allow()
}

func main() {
	ul := NewUserLimiter()

	user := "user1"

	for i := 0; i < 18; i++ {
		if ul.Allow(user) {
			fmt.Println("Allowed", i)
		} else {
			fmt.Println("Blocked", i)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
