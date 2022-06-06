// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validations

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"phantomDNS/repositories"
)

func DnsQueryAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		clientIP := context.GetHeader("X-Forwarded-For")
		fmt.Print("X-Forwarded-For==================" + clientIP + "\n")
		client := repositories.QueryClientByIp(clientIP)
		if client.ID > 0 {
			context.Next()
		} else {
			context.Abort()
			context.String(http.StatusNotFound, "404 page not found")
		}
	}
}
