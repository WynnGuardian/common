package middlewares

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wynnguardian/common/handlerfunc"
	"github.com/wynnguardian/common/response"
	"golang.org/x/time/rate"
)

func RateLimit(rl *RateLimiter, whitelist []string, next handlerfunc.HandlerFunc) handlerfunc.HandlerFunc {
	return func(ctx *gin.Context) response.WGResponse {

		for _, w := range whitelist {
			if w == ctx.GetHeader("Authorization") {
				return next(ctx)
			}
		}

		clientIP := ctx.ClientIP()

		limiter := rl.GetLimiter(clientIP)

		if !limiter.Allow() {
			return response.ErrTooManyRequests
		}

		return next(ctx)
	}
}

type RateLimiter struct {
	ipLimiters map[string]*rate.Limiter
	mu         sync.Mutex
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		ipLimiters: make(map[string]*rate.Limiter),
	}
}

func (rl *RateLimiter) GetLimiter(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	limiter, exists := rl.ipLimiters[ip]
	if !exists {
		limiter = rate.NewLimiter(rate.Every(2*time.Second), 1)
		rl.ipLimiters[ip] = limiter
	}
	return limiter
}
