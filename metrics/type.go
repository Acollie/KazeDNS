package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	OpsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "dns_processed_ops_total",
		Help: "The total number of processed events",
	})
	CacheHits = promauto.NewCounter(prometheus.CounterOpts{
		Name: "dns_cache_hits_total",
		Help: "The total number of cache hits",
	})
	CacheMisses = promauto.NewCounter(prometheus.CounterOpts{
		Name: "dns_cache_misses_total",
		Help: "The total number of cache misses",
	})
	Blocked = promauto.NewCounter(prometheus.CounterOpts{
		Name: "dns_blocked_total",
		Help: "The total number of blocked requests",
	})
	Failures = promauto.NewCounter(prometheus.CounterOpts{
		Name: "dns_failures_total",
		Help: "The total number of failures",
	})
)
