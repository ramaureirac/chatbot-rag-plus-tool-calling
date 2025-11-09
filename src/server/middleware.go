package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func xAnonIDMiddleware(sess *Sessions) gin.HandlerFunc {
	return func(c *gin.Context) {
		anonID := c.GetHeader("X-Anon-ID")
		sess.Mutex.Lock()
		_, ok := sess.Agents[anonID]
		sess.Mutex.Unlock()
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "unauthenticated",
			})
			c.Abort()
		}
		c.Next()
	}
}
