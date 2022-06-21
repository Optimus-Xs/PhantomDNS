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

// serve init steps, read publicDNS from app.yml
func init() {
	utils.ReadConfig()
	publicDNS = viper.GetString("publicDNS")
}

// DnsServicesInterfaces define Dns related methods
type DnsServicesInterfaces interface {
	DNSQuery(dnsQuery []byte) (httpCode int, res []byte)                              // DoH query use binary dns message
	DNSResolve(domain string, dnsType uint16) (httpCode int, res entities.JsonDnsMsg) // DoH query use json dns message
}

// DnsServices is implantation of DnsServicesInterfaces
type DnsServices struct {
}

// DNSQuery is implantation of DnsServicesInterfaces.DNSQuery it gets binary dns message and match custom domain in db
// if it finds a matched domain it will pack a binary answer dns message and return, else query publicDNS use binary dns
// message and return resolve result from publicDNS
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

// DNSResolve is implantation of DnsServicesInterfaces.DNSResolve it gets json dns message and match custom domain in db
// if it finds a matched domain it will return resolve result in json format, else query publicDNS with json and return
// resolve result from publicDNS
func (service DnsServices) DNSResolve(domain string, dnsType uint16) (httpCode int, res entities.JsonDnsMsg) {
	dnsRecord := repositories.QueryDnsByDomain(domain)
	if dnsRecord.ID > 0 {
		res = utils.DnsResJsonBuilder(dnsRecord, domain, dnsType)
		return 200, res
	} else {
		return jsonPubQuery(domain, dnsType)
	}

}

// dnsMsgPubQuery provide service to query publicDNS by use binary dns message
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

// jsonPubQuery provide service to query publicDNS by use json dns message
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

// requestResolve can unpack binary dns message into dns.Msg which allowed to easy access info inside the DNS message
func requestResolve(dnsRequestMessage []byte) (dnsMsg dns.Msg) {
	dnsMsg = dns.Msg{}
	dnsMsg.Unpack(dnsRequestMessage)
	return dnsMsg
}
