package middleware

import (
	"github.com/gin-gonic/gin"
	wework "github.com/haiyiyun/weworksdk"
)

func InjectSdk(ww wework.IWeWork) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("ww", ww)
		c.Next()
	}
}
