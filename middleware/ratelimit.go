package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := NewTokenBucket(int64(fillInterval.Seconds()), cap)
	return func(c *gin.Context) {
		if !bucket.Allow() {
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}

type TokenBucket struct {
	rate         int64
	capacity     int64
	tokenSize    int64
	laskTokenSec int64
	lock         sync.Mutex
}

func NewTokenBucket(rate, cap int64) *TokenBucket {
	return &TokenBucket{
		rate:         rate,
		capacity:     cap,
		tokenSize:    0,
		laskTokenSec: time.Now().Unix(),
	}
}

func (b *TokenBucket) Set(rate, cap int64) {
	b.rate = rate
	b.capacity = cap
	b.tokenSize = 0
	b.laskTokenSec = time.Now().Unix()
}

func (b *TokenBucket) Allow() bool {
	b.lock.Lock()
	defer b.lock.Unlock()
	now := time.Now().Unix()
	b.tokenSize = b.tokenSize + (now-b.laskTokenSec)*b.rate

	if b.tokenSize > b.capacity {
		b.tokenSize = b.capacity
	}
	b.laskTokenSec = now
	if b.tokenSize > 0 {
		b.tokenSize--
		return true
	} else {
		return false
	}

}
