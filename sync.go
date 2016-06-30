package bitters

import (
	"github.com/gin-gonic/gin"
	"sync"
)

func Sync(wg *sync.WaitGroup) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer wg.Done()
		wg.Add(1)
		c.Next()
	}
}
