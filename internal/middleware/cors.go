package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type CORS struct {
	allowed []string
}

func NewCORS(origins []string) *CORS { return &CORS{allowed: origins} }

func (m *CORS) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		allow := ""
		if len(m.allowed) == 1 && m.allowed[0] == "*" {
			allow = "*"
		} else {
			for _, o := range m.allowed {
				if strings.EqualFold(o, origin) {
					allow = origin
					break
				}
			}
		}
		if allow != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", allow)
		}
		c.Writer.Header().Set("Vary", "Origin")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
