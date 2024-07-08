package resolver

import (
	"github.com/miekg/dns"
)

type Handler struct {
	cache     dnsCache
	blocklist blockList
}
type blockList map[string]bool
type dnsCache map[string]*dns.Msg

func New() *Handler {
	return &Handler{
		cache:     make(dnsCache),
		blocklist: make(blockList),
	}
}
