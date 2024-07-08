package main

import (
	"dns-server/resolver"
	"flag"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/miekg/dns"
)

func main() {
	port := 8888
	flag.Parse()

	handler := resolver.New()

	for z := range resolver.Zones {
		z := z
		dns.HandleFunc(z, func(writer dns.ResponseWriter, msg *dns.Msg) {
			handler.DNSResolver(writer, msg)
		})
	}

	dns.HandleFunc(".", func(writer dns.ResponseWriter, msg *dns.Msg) {
		handler.DNSResolver(writer, msg)
	})

	go func() {
		srv := &dns.Server{Addr: ":" + strconv.Itoa(port), Net: "udp"}
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Failed to set udp listener %s\n", err.Error())
		}
	}()

	go func() {
		srv := &dns.Server{Addr: ":" + strconv.Itoa(port), Net: "tcp"}
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Failed to set tcp listener %s\n", err.Error())
		}
	}()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig
	log.Fatalf("Signal (%v) received, stopping\n", s)
}
