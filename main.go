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

	config := &dns.ClientConfig{
		Servers: []string{resolver.DnsGoogleServer, resolver.DnsCloudFlare},
	}

	handler := resolver.New(config)

	for z := range resolver.Zones {
		z := z
		dns.HandleFunc(z, func(writer dns.ResponseWriter, msg *dns.Msg) {
			err := handler.DNSResolver(writer, msg)
			if err != nil {
				log.Printf("Failed to resolve %s: %v", z, err)
			}
		})
	}

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
