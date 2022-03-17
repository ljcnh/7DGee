/**
 * @Author: lj
 * @Description:
 * @File:  logger
 * @Version: 1.0.0
 * @Date: 2022/03/17 15:29
 */

package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(context *Context) {
		t := time.Now()
		context.Next()
		log.Printf("[%d] %s in %v", context.StatusCode, context.Req.RequestURI, time.Since(t))
	}
}
