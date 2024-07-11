package resolver

import (
	"fmt"
	"github.com/miekg/dns"
	"log"
)

const (
	dnsPort   = 53
	dnsServer = "8.8.8.8"
)

func (h *Handler) DNSResolver(w dns.ResponseWriter, r *dns.Msg) {
	c := new(dns.Client)
	m := new(dns.Msg)

	if len(r.Question) == 0 {
		log.Println("No questions in the request")
		return
	}

	for _, question := range r.Question {
		err := h.blocklist.BlockItems.Check(question.Name)
		if err != nil {
			return
		}
	}

	in, _, err := c.Exchange(m, fmt.Sprintf("%s:%d", dnsServer, dnsPort))
	if err != nil {
		log.Printf("Failed to query root DNS server: %v", err)
		return
	}
	response := new(dns.Msg)
	response.SetReply(r)
	response.Authoritative = true
	for _, ans := range in.Answer {
		if aRecord, ok := ans.(*dns.A); ok {
			response.Answer = append(response.Answer, aRecord)
		}
	}
	err = w.WriteMsg(response)
	if err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
