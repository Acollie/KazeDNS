package resolver

import (
	"dns-server/blocklist"
	"github.com/miekg/dns"
)

type Handler struct {
	cache     dnsCache
	blocklist *blocklist.BlocksCli
}
type dnsCache map[string]*dns.Msg

func New() *Handler {
	return &Handler{
		cache:     make(dnsCache),
		blocklist: blocklist.New(),
	}
}
