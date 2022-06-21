// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package utils

import (
	"github.com/miekg/dns"
	"github.com/spf13/viper"
	"net"
	"phantomDNS/entities"
	"strconv"
	"strings"
)

var ttl int

// read DNS answer record ttl from app.yml
func init() {
	ReadConfig()
	ttl = viper.GetInt("DDns.ttl")
}

// DnsResMsgBuilder can build DNS answer message from dnsRecord and query DNS message
func DnsResMsgBuilder(dnsRecord entities.DnsRecord, dnsQueryMsg dns.Msg) (res dns.Msg) {
	resourceDataLength := 4
	if dnsRecord.DnsType == dns.TypeAAAA {
		resourceDataLength = 16
	}

	hdr := dns.MsgHdr{
		Id:                 dnsQueryMsg.MsgHdr.Id,
		Response:           true,
		Opcode:             0,
		Authoritative:      false,
		Truncated:          false,
		RecursionDesired:   true,
		RecursionAvailable: true,
		Zero:               false,
		AuthenticatedData:  false,
		CheckingDisabled:   false,
		Rcode:              0,
	}
	ans := dns.A{
		Hdr: dns.RR_Header{
			Name:     dnsRecord.Domain,
			Rrtype:   dnsQueryMsg.Question[0].Qtype,
			Class:    1,
			Ttl:      uint32(ttl),
			Rdlength: uint16(resourceDataLength),
		},
		A: net.ParseIP(dnsRecord.IpAddress),
	}

	res.MsgHdr = hdr
	res.Question = dnsQueryMsg.Question
	res.Answer = append(res.Answer, &ans)
	return res
}

// DnsResJsonBuilder can build DNS answer message in json format from dnsRecord, query domain and query Dns type
func DnsResJsonBuilder(dnsRecord entities.DnsRecord, domain string, dnsType uint16) (res entities.JsonDnsMsg) {
	res = entities.JsonDnsMsg{
		Status:   0,
		TC:       false,
		RD:       true,
		RA:       true,
		AD:       false,
		CD:       false,
		Question: nil,
		Answer:   nil,
	}

	res.Question = append(res.Question, entities.Question{
		Name: domain,
		Type: dnsType,
	})

	res.Answer = append(res.Answer, entities.Answer{
		Name: dnsRecord.Domain,
		Type: dnsRecord.DnsType,
		TTL:  ttl,
		Data: dnsRecord.IpAddress,
	})

	return res
}

// DnsTypeConverter convert DNS record type from string to uint16
func DnsTypeConverter(dnsType string) (convertedType uint16) {
	IntDnsType, err := strconv.ParseInt(dnsType, 10, 16)
	if err == nil {
		return uint16(IntDnsType)
	}
	return dnsTypeMapping()(dnsType)
}

// dnsTypeMapping is mapping rule of DNS record type between sting and uint16
func dnsTypeMapping() func(string) uint16 {
	// typeMapping is captured in the closure returned below
	typeMapping := map[string]uint16{
		"None":       0,
		"A":          1,
		"NS":         2,
		"MD":         3,
		"MF":         4,
		"CNAME":      5,
		"SOA":        6,
		"MB":         7,
		"MG":         8,
		"MR":         9,
		"NULL":       10,
		"PTR":        12,
		"HINFO":      13,
		"MINFO":      14,
		"MX":         15,
		"TXT":        16,
		"RP":         17,
		"AFSDB":      18,
		"X25":        19,
		"ISDN":       20,
		"RT":         21,
		"NSAPPTR":    23,
		"SIG":        24,
		"KEY":        25,
		"PX":         26,
		"GPOS":       27,
		"AAAA":       28,
		"LOC":        29,
		"NXT":        30,
		"EID":        31,
		"NIMLOC":     32,
		"SRV":        33,
		"ATMA":       34,
		"NAPTR":      35,
		"KX":         36,
		"CERT":       37,
		"DNAME":      39,
		"OPT":        41,
		"APL":        42,
		"DS":         43,
		"SSHFP":      44,
		"RRSIG":      46,
		"NSEC":       47,
		"DNSKEY":     48,
		"DHCID":      49,
		"NSEC3":      50,
		"NSEC3PARAM": 51,
		"TLSA":       52,
		"SMIMEA":     53,
		"HIP":        55,
		"NINFO":      56,
		"RKEY":       57,
		"TALINK":     58,
		"CDS":        59,
		"CDNSKEY":    60,
		"OPENPGPKEY": 61,
		"CSYNC":      62,
		"ZONEMD":     63,
		"SVCB":       64,
		"HTTPS":      65,
		"SPF":        99,
		"UINFO":      100,
		"UID":        101,
		"GID":        102,
		"UNSPEC":     103,
		"NID":        104,
		"L32":        105,
		"L64":        106,
		"LP":         107,
		"EUI48":      108,
		"EUI64":      109,
		"URI":        256,
		"CAA":        257,
		"AVC":        258,
		"TKEY":       249,
		"TSIG":       250,
		"IXFR":       251,
		"AXFR":       252,
		"MAILB":      253,
		"MAILA":      254,
		"ANY":        255,
		"TA":         32768,
		"DLV":        32769,
		"Reserved":   65535,
	}

	return func(key string) uint16 {
		return typeMapping[key]
	}
}

// GetIpType can identify string ip address is ipv4 or ipv6
func GetIpType(ipAddress string) uint16 {
	if strings.Count(ipAddress, ":") >= 2 {
		return dns.TypeAAAA
	} else {
		return dns.TypeA
	}
}
