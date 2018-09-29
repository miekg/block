package block

import (
	"sync"

	"github.com/coredns/coredns/plugin"

	"github.com/prometheus/client_golang/prometheus"
)

var blockCount = prometheus.NewCounterVec(prometheus.CounterOpts{
	Namespace: plugin.Namespace,
	Subsystem: "block",
	Name:      "count_total",
	Help:      "Counter of blocked names.",
}, []string{"server"})

var once sync.Once
