package resolver

import (
	"dns-server/blocklist"
	"github.com/miekg/dns"
)

type Handler struct {
	cache     dnsCache
	config    *dns.ClientConfig
	blocklist *blocklist.BlocksCli
}
type dnsCache map[string]*dns.Msg

func New(config *dns.ClientConfig) *Handler {
	if config == nil {
		config = &dns.ClientConfig{
			Servers: []string{DnsGoogleServer, DnsCloudFlare},
		}
	}
	return &Handler{
		cache:     make(dnsCache),
		blocklist: blocklist.New(),
		config:    config,
	}
}
