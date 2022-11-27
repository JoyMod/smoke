package smoke

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		//开始计时
		t := time.Now()

		c.Next()

		//日志打印
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
