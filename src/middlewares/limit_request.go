package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Client struct {
	Requests int
	LastTime time.Time
}


var rateLimit = 40
var timeWindow = time.Minute
var clients = make(map[string]*Client)
var mu sync.Mutex




func RateLimitMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        clientIP := c.ClientIP()

        mu.Lock()
        client, exists := clients[clientIP]
        if !exists {
            client = &Client{Requests: 1, LastTime: time.Now()}
            clients[clientIP] = client
        } else {
            elapsed := time.Since(client.LastTime)
            if elapsed > timeWindow {
                client.Requests = 1
                client.LastTime = time.Now()
            } else {
                client.Requests++
            }
        }
        mu.Unlock()
        if client.Requests > rateLimit {
            c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
            c.Abort()
            return
        }

        c.Next()
    }
}









