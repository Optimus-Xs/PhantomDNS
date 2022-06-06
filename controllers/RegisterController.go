// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"phantomDNS/services"
	"phantomDNS/utils"
)

func SetRegistryController(router *gin.Engine) {
	//router.Use(gin.BasicAuth(services.RegisterServices{}.GetDevices()))

	router.POST("/registerHost", gin.BasicAuth(services.RegisterServices{}.GetDevices()), func(context *gin.Context) {
		ip, domain := context.Query("ip"), context.Query("domain")
		services.RegisterServices{}.DoHostRegister(ip, domain, utils.GetCurrentAuthDev(context))
		context.String(http.StatusOK, "success")
	})

	router.POST("/registerClient", gin.BasicAuth(services.RegisterServices{}.GetDevices()), func(context *gin.Context) {
		ip := context.Query("ip")
		services.RegisterServices{}.DoClientRegister(ip, utils.GetCurrentAuthDev(context))
		context.String(http.StatusOK, "success")
	})
}
