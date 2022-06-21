// Copyright [2022] [Optimus-Xs@GitHub]. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package entities

// JsonDnsMsg is object to handle DNS over HTTPS query in json format, which contained json struct for DOH query and answer
type JsonDnsMsg struct {
	Status   int        `json:"Status"`   // Standard DNS response code (32 bit integer).
	TC       bool       `json:"TC"`       // Whether the response is truncated
	RD       bool       `json:"RD"`       // Always true for Google Public DNS
	RA       bool       `json:"RA"`       // Always true for Google Public DNS
	AD       bool       `json:"AD"`       // Whether all response data was validated with DNSSEC
	CD       bool       `json:"CD"`       // Whether the client asked to disable DNSSEC
	Question []Question `json:"Question"` // DNS query info
	Answer   []Answer   `json:"Answer"`   // DNS answer info
}

// Question is query info part of DOH massage in json format
type Question struct {
	Name string `json:"name"` // FQDN with trailing dot
	Type uint16 `json:"type"` // A - Standard DNS RR type
}

// Answer is answer info part of DOH massage in json format
type Answer struct {
	Name string `json:"name"` // Always matches name in the Question section
	Type uint16 `json:"type"` // A - Standard DNS RR type
	TTL  int    `json:"TTL"`  // Record's time-to-live in seconds
	Data string `json:"data"` // Data for A - IP address as text
}
