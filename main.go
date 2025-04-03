package main

import (
	"dns-server/blocklist"
	"dns-server/resolver"
	"flag"
	"github.com/miekg/dns"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	port = 8888
)

func main() {
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

	defaultList := blocklist.LoadBlockList("https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts")
	handler.AddBatch(defaultList)

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/remove", http.HandlerFunc(handler.Remove))
	http.Handle("/add", http.HandlerFunc(handler.AddHttp))
	http.Handle("/list", http.HandlerFunc(handler.Get))
	http.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
		return
	}))
	http.ListenAndServe(":2112", nil)

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig
	log.Fatalf("Signal (%v) received, stopping\n", s)
}
