package resolver

import (
	"github.com/miekg/dns"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_DNSResolver(t *testing.T) {

	t.Run("No questions in the request", func(t *testing.T) {
		handler := Handler{}
		resonse := dns.ResponseWriter(nil)
		msg := new(dns.Msg)
		err := handler.DNSResolver(resonse, msg)
		require.Error(t, err)
	})
}
