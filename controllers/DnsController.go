// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controllers

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"phantomDNS/services"
	"phantomDNS/services/validations"
	"phantomDNS/utils"
)

func SetDnsController(router *gin.Engine) {
	router.GET("/dns-query", validations.DnsQueryAuth(), func(context *gin.Context) {
		dnsQueryBase64 := context.Query("dns")
		dnsQuery, _ := base64.RawURLEncoding.DecodeString(dnsQueryBase64)
		httpCode, queryRes := services.DnsServices{}.DNSQuery(dnsQuery)
		context.Data(httpCode, "application/dns-message", queryRes)
	})

	router.POST("/dns-query", validations.DnsQueryAuth(), func(context *gin.Context) {
		dnsQuery, _ := ioutil.ReadAll(context.Request.Body)
		httpCode, queryRes := services.DnsServices{}.DNSQuery(dnsQuery)
		context.Data(httpCode, "application/dns-message", queryRes)
	})

	router.GET("/resolve", validations.DnsQueryAuth(), func(context *gin.Context) {
		domain := context.Query("name")
		dnsTypeSting := context.Query("type")
		dnsType := utils.DnsTypeConverter(dnsTypeSting)
		httpCode, queryRes := services.DnsServices{}.DNSResolve(domain, dnsType)
		context.JSON(httpCode, queryRes)
	})
}
