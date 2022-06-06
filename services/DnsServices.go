// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package services

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/miekg/dns"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"os"
	"phantomDNS/entities"
	"phantomDNS/repositories"
	"phantomDNS/utils"
	"strconv"
)

var (
	publicDNS string
)

func init() {
	utils.ReadConfig()
	publicDNS = viper.GetString("publicDNS")
}

type DnsServicesInterfaces interface {
	DNSQuery(dnsQuery []byte) (httpCode int, res []byte)
	DNSResolve(domain string, dnsType uint16) (httpCode int, res entities.JsonDnsMsg)
}

type DnsServices struct {
}

func (service DnsServices) DNSQuery(dnsQuery []byte) (httpCode int, res []byte) {
	dnsQueryMsg := requestResolve(dnsQuery)
	domain := dnsQueryMsg.Question[0].Name
	fmt.Printf("Query Domian： " + domain + "\n")
	dnsRecord := repositories.QueryDnsByDomain(domain)
	if dnsRecord.ID > 0 {
		dDnsRes := utils.DnsResMsgBuilder(dnsRecord, dnsQueryMsg)
		dnsMsg, _ := dDnsRes.Pack()
		fmt.Printf("DDNS Query Result： " + dnsRecord.IpAddress + "\n")
		return 200, dnsMsg
	} else {
		return dnsMsgPubQuery(dnsQuery)
	}
}

func (service DnsServices) DNSResolve(domain string, dnsType uint16) (httpCode int, res entities.JsonDnsMsg) {
	dnsRecord := repositories.QueryDnsByDomain(domain)
	if dnsRecord.ID > 0 {
		res = utils.DnsResJsonBuilder(dnsRecord, domain, dnsType)
		return 200, res
	} else {
		return jsonPubQuery(domain, dnsType)
	}

}

func dnsMsgPubQuery(dnsQuery []byte) (httpCode int, queryRes []byte) {
	b64 := base64.RawURLEncoding.EncodeToString(dnsQuery)
	url := "https://" + publicDNS + "/dns-query?dns=" + b64
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Send query error, err:%v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	httpCode = resp.StatusCode
	queryRes, _ = ioutil.ReadAll(resp.Body)
	return httpCode, queryRes
}

func jsonPubQuery(domain string, dnsType uint16) (httpCode int, queryRes entities.JsonDnsMsg) {
	resp, err := http.Get("https://" + publicDNS + "/resolve?name=" + domain + "&type=" + strconv.Itoa(int(dnsType)))
	if err != nil {
		fmt.Printf("Send query error, err:%v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	httpCode = resp.StatusCode
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &queryRes)
	return httpCode, queryRes
}

func requestResolve(dnsRequestMessage []byte) (dnsMsg dns.Msg) {
	dnsMsg = dns.Msg{}
	dnsMsg.Unpack(dnsRequestMessage)
	return dnsMsg
}
