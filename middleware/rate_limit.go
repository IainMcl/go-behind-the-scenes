package middleware

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
)

// TODO: This is incomplete but not needed
func RateLimiter(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.MustGet("userRole").(string) // get user role from context

		// Set different rate limits based on user role
		switch userRole {
		case "admin":
			lmt.SetMax(5) // 5 requests per second for admin
		case "user":
			lmt.SetMax(1) // 1 request per second for user
		default:
			lmt.SetMax(0.5) // 0.5 requests per second for others
		}

		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.Data(httpError.StatusCode, c.ContentType(), []byte(httpError.Message))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
