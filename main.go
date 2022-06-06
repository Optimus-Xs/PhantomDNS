// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gin-gonic/gin"
	"phantomDNS/controllers"
)

var (
	r = gin.Default()
)

func init() {
	controllers.SetDnsController(r)
	controllers.SetRegistryController(r)
}

func main() {

	r.Run(":8000")
}
