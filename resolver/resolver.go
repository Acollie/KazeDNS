package resolver

import (
	"errors"
	"fmt"
	"github.com/miekg/dns"
	"log"
)

const (
	dnsPort         = 53
	DnsCloudFlare   = "1.1.1.1"
	DnsGoogleServer = "8.8.8.8"
)

func (h *Handler) DNSResolver(w dns.ResponseWriter, r *dns.Msg) error {
	c := new(dns.Client)
	m := new(dns.Msg)

	if len(r.Question) == 0 {
		log.Println("No questions in the request")
		return errors.New("no questions in the request")
	}

	for _, question := range r.Question {
		err := h.blocklist.BlockItems.Check(question.Name)
		if err != nil {
			return fmt.Errorf("blocked: %s", question.Name)
		}
	}

	var in *dns.Msg
	var err error
	for _, server := range h.config.Servers {
		in, _, err = c.Exchange(m, fmt.Sprintf("%s:%d", server, dnsPort))
		if err != nil {
			log.Printf("Failed to query root DNS server: %v", err)
			continue
		}
		break
	}

	if in == nil {
		return errors.New("all DNS queries failed")
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
		return fmt.Errorf("failed to write response: %v", err)
	}
	return nil
}
