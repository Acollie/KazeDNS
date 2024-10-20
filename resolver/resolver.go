package resolver

import (
	"dns-server/metrics"
	"errors"
	"fmt"
	"github.com/miekg/dns"
	"log"
)

func (h *Handler) DNSResolver(w dns.ResponseWriter, r *dns.Msg) error {

	metrics.OpsProcessed.Inc()
	c := new(dns.Client)
	m := new(dns.Msg)

	if len(r.Question) == 0 {
		metrics.Failures.Inc()
		return errors.New("no questions in the request")
	}

	for _, question := range r.Question {
		err := h.blocklist.BlockItems.Check(question.Name)
		if err != nil {
			metrics.Blocked.Inc()
			return fmt.Errorf("blocked: %s", question.Name)
		}
	}

	var in *dns.Msg
	var err error
	for _, server := range h.config.Servers {
		in, _, err = c.Exchange(m, fmt.Sprintf("%s:%d", server, dnsPort))
		if err != nil {
			metrics.Failures.Inc()
			log.Printf("Failed to query root DNS server: %v", err)
			continue
		}
		break
	}

	if in == nil {
		metrics.Failures.Inc()
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
